package service

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/MisakiFx/martin/pkg/constant"
	"github.com/MisakiFx/martin/pkg/dao"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/MisakiFx/martin/pkg/tools"
)

func BookingCheckService(req *model.BookingCheckReq, openId string) (int64, int, error) {
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
	allPay *= examination.UserCardType / 10
	lastRemainder := examination.UserRemainder
	lastCheckCount := examination.UserCheckCount
	if req.PayType == constant.PayTypeRemainder {
		if lastRemainder < allPay {
			tx.Rollback()
			tools.GetLogger().Errorf("余额不足,请充值")
			return 0, constant.StatusCodeInputError, errors.New("余额不足,请充值")
		}
		lastRemainder -= allPay
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
	bookingId := tools.GenId()
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
