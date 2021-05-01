package service

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/dao"
	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/MisakiFx/martin/martin/pkg/tools"
)

func BookingCheckService(req *model.BookingCheckReq, openId string) (int64, int, error) {
	sort.Ints(req.CheckProject)
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return 0, constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	startTime, err := time.ParseInLocation(constant.TimeFormatString, req.StartTime, tools.LocGloble)
	if err != nil {
		tools.GetLogger().Errorf("service.BookingCheckService parse time error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	var allPay float64
	var projectsString string
	for i, pro := range req.CheckProject {
		allPay += model.CheckProjectMap[pro].Money
		if i == 0 {
			projectsString += strconv.FormatInt(int64(pro), 10)
			continue
		}
		projectsString += "," + strconv.FormatInt(int64(pro), 10)
	}

	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)

	//校验是否存在未结束的预约
	bookingInfo, err := dao.GetStartedBookingCheck(tx, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.GetStartedBookingCheck error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if bookingInfo != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.GetStartedBookingCheck has booked")
		return 0, constant.StatusCodeInputError, errors.New("存在未结束的体检预约,不能同时预约多次体检")
	}

	//校验预约时段是否已满
	bookedCount, err := dao.CountBookedCheckTime(tx, startTime)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.CountBookedCheckTime error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if bookedCount > constant.BookCheckMaxCount {
		tx.Rollback()
		errMsg := "当前时间段预约已满"
		tools.GetLogger().Errorf(errMsg)
		return 0, constant.StatusCodeInputError, errors.New(errMsg)
	}

	//校验余额是否充足
	examination, err := dao.GetUserExamination(tx, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.GetUserExamination error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if examination == nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.GetUserExamination error : %v", "找不到用户余额信息")
		return 0, constant.StatusCodeServiceError, errors.New("找不到用户余额信息")
	}
	allPay = allPay * examination.UserCardType / 10
	allPay = tools.FloatRound(allPay, 2)
	lastRemainder := examination.UserRemainder
	lastCheckCount := examination.UserCheckCount
	if req.PayType == constant.PayTypeRemainder {
		if lastRemainder < allPay {
			tx.Rollback()
			tools.GetLogger().Errorf("余额不足,请充值")
			return 0, constant.StatusCodeInputError, errors.New("余额不足,请充值")
		}
		lastRemainder -= allPay
		lastRemainder = tools.FloatRound(lastRemainder, 2)
	}
	if req.PayType == constant.PayTypeCheckCount {
		if lastCheckCount < 1 {
			tx.Rollback()
			tools.GetLogger().Errorf("余额不足,请充值")
			return 0, constant.StatusCodeInputError, errors.New("余额不足,请充值")
		}
		lastCheckCount -= 1
	}

	//更新余额信息
	effectRows, err := dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: lastCheckCount,
		UserRemainder:  lastRemainder,
		UserCardType:   examination.UserCardType,
		UpdateTime:     examination.UpdateTime,
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.UpdateUserExamination error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if effectRows == 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.UpdateUserExamination effect rows = 0")
		return 0, constant.StatusCodeInputError, errors.New("操作太频繁,请稍后重试")
	}
	//插入预约表
	var payStruct model.BookingPay
	if req.PayType == constant.PayTypeRemainder {
		payStruct = model.BookingPay{
			Remainder:  allPay,
			CheckCount: 0,
		}
	} else {
		payStruct = model.BookingPay{
			Remainder:  0,
			CheckCount: 1,
		}
	}
	payString, err := json.Marshal(&payStruct)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService marshal pay string error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	bookingId := tools.GenId()
	err = dao.CreateBookingCheck(tx, &model.GuardianBookingInfo{
		ID:           bookingId,
		UserId:       userInfo.ID,
		CheckProject: projectsString,
		StartTime:    startTime,
		EndTime:      startTime.Add(time.Hour * 2),
		Status:       0,
		Pay:          string(payString),
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	})
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.CreateBookingCheck error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	//插入结果表
	err = dao.CreateCheckResult(tx, &model.GuardianCheckResult{
		ID:            tools.GenId(),
		BookingId:     bookingId,
		UserId:        userInfo.ID,
		Internal:      "",
		Surgery:       "",
		Sgpt:          "",
		BloodGlucode:  "",
		BloodFat:      "",
		RenalFunction: "",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	})
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BookingCheckService->dao.CreateCheckResult error : %v", err)
		return 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	tx.Commit()

	//todo 起一个timer定时提醒
	return bookingId, constant.StatusCodeSuccess, nil
}

func ListCheckService(openId string, page, size int) (int64, []model.ListCheckResp, int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return 0, nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return 0, nil, constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	count, checks, err := dao.ListCheck(userInfo.ID, page, size)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.ListCheckService error : %v", err)
		return 0, nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	list := make([]model.ListCheckResp, 0)
	for _, check := range checks {
		proString := strings.Split(check.CheckProject, ",")
		proInt := make([]int, 0)
		for _, pro := range proString {
			proIntTemp, err := strconv.ParseInt(pro, 10, 32)
			if err != nil {
				tools.GetLogger().Errorf("service.RefundExaminationService : 存在不合法的检查项目")
				return 0, nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
			}
			proInt = append(proInt, int(proIntTemp))
		}
		var pay model.BookingPay
		err = json.Unmarshal([]byte(check.Pay), &pay)
		if err != nil {
			tools.GetLogger().Errorf("service.RefundExaminationService : 付款方式不合法")
			return 0, nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
		}
		list = append(list, model.ListCheckResp{
			Id:            check.ID,
			CheckProject:  proInt,
			StartTime:     check.StartTime.Format(constant.TimeFormatString),
			EndTime:       check.EndTime.Format(constant.TimeFormatString),
			Status:        check.Status,
			CreateTime:    check.CreateTime.Format(constant.TimeFormatString),
			PayReminder:   pay.Remainder,
			PayCheckCount: pay.CheckCount,
		})
	}
	return count, list, constant.StatusCodeSuccess, nil
}

func GetCheckResultService(openId string, bookingId int64) (*model.GetCheckResultResp, int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.GetCheckResultService->dao.GetUserInfoByOpenId error : %v", err)
		return nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.GetCheckResultService user not found")
		return nil, constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	checkInfo, err := dao.GetCheckInfo(bookingId, userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.GetCheckResultService->dao.GetCheckInfo error : %v", err)
		return nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if checkInfo == nil {
		tools.GetLogger().Warn("service.GetCheckResultService->dao.GetCheckInfo can not found id")
		return &model.GetCheckResultResp{}, constant.StatusCodeInputError, errors.New("没有权限查看该体检结果")
	}
	projectsInt := make([]int, 0)
	projectsString := strings.Split(checkInfo.CheckProject, ",")
	for _, projectString := range projectsString {
		proTemp, _ := strconv.ParseInt(projectString, 10, 64)
		projectsInt = append(projectsInt, int(proTemp))
	}

	result, err := dao.GetCheckResultByBookingId(bookingId, userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.GetCheckResultService->dao.GetCheckResultByBookingId error : %v", err)
		return nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if result == nil {
		tools.GetLogger().Warn("service.GetCheckResultService->dao.GetCheckResultByBookingId can not found id")
		return &model.GetCheckResultResp{}, constant.StatusCodeInputError, errors.New("没有权限查看该体检结果")
	}
	return &model.GetCheckResultResp{
		BookingId:     result.BookingId,
		Projects:      projectsInt,
		Internal:      result.Internal,
		Surgery:       result.Surgery,
		ENT:           result.Ent,
		SGPT:          result.Sgpt,
		BloodGlucode:  result.BloodGlucode,
		BloodFat:      result.BloodFat,
		RenalFunction: result.RenalFunction,
	}, constant.StatusCodeSuccess, nil
}

func CancelBookingCheckService(openId string, bookingId int64) (int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.GetUserInfoByOpenId error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.CancelBookingCheckService user not found")
		return constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}

	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	//获取体检信息
	bookingInfo, err := dao.GetUserBookingInfo(userInfo.ID, bookingId)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.GetUserBookingInfo error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	if bookingInfo == nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService booking is not belong to user")
		return constant.StatusCodeInputError, errors.New("无权限删除该体检")
	}

	//验证
	if bookingInfo.Status != 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService check is start")
		return constant.StatusCodeInputError, errors.New("体检已经开始不能取消")
	}
	var pay model.BookingPay
	err = json.Unmarshal([]byte(bookingInfo.Pay), &pay)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService : unmarshal booking pay error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}

	//删除体检记录
	rows, err := dao.DeleteCheck(tx, userInfo.ID, bookingId)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.DeleteCheck error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	if rows == 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.DeleteCheck rows effect = 0")
		return constant.StatusCodeInputError, errors.New("未找到对应体检记录")
	}

	//退款
	examination, err := dao.GetUserExamination(tx, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.GetUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if examination == nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.GetUserExamination error : %v", "找不到用户余额信息")
		return constant.StatusCodeServiceError, errors.New("找不到用户余额信息")
	}

	effectRows, err := dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: examination.UserCheckCount + pay.CheckCount,
		UserRemainder:  tools.FloatRound(examination.UserRemainder+pay.Remainder, 2),
		UserCardType:   examination.UserCardType,
		UpdateTime:     examination.UpdateTime,
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.UpdateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if effectRows == 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.CancelBookingCheckService->dao.UpdateUserExamination effect rows = 0")
		return constant.StatusCodeInputError, errors.New("操作太频繁,请稍后重试")
	}
	tx.Commit()
	return constant.StatusCodeSuccess, nil
}
