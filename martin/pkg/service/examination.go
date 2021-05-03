package service

import (
	"errors"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/pkg/dao"
	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/MisakiFx/martin/martin/pkg/tools"
)

func BuyExaminationService(req *model.BuyExaminationReq, openId string) (int, error) {
	examination := model.ExaminationMap[req.ExaminationId]
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	examinationOld, err := dao.GetUserExamination(nil, userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.GetUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}

	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	//更新用户体检卡信息
	cardType := examinationOld.UserCardType
	if examination.CardType < cardType {
		cardType = examination.CardType
	}
	effectRows, err := dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: examinationOld.UserCheckCount + examination.CheckCount,
		UserRemainder:  tools.FloatRound(examinationOld.UserRemainder+examination.Remainder, 2),
		UserCardType:   cardType,
		UpdateTime:     examinationOld.UpdateTime,
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.CreateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if effectRows == 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.CreateUserExamination effect rows = 0")
		return constant.StatusCodeInputError, errors.New("操作太频繁,请稍后重试")
	}
	//插入消费记录
	err = dao.CreateExpenseCalendar(tx, &model.GuardianExpenseCalendar{
		ID:         tools.GenId(),
		UserId:     userInfo.ID,
		Money:      examination.Cost,
		Status:     constant.ExpenseStatusCost,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.CreateExpenseCalendar error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	tx.Commit()
	return constant.StatusCodeSuccess, nil
}

func GetExaminationInfoService(openId string) (*model.GetExaminationInfoResp, int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return nil, constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	examination, err := dao.GetUserExamination(nil, userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.GetUserExamination error : %v", err)
		return nil, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}

	return &model.GetExaminationInfoResp{
		CheckCount: examination.UserCheckCount,
		Remainder:  examination.UserRemainder,
		CardType:   examination.UserCardType,
	}, constant.StatusCodeSuccess, nil
}

func RefundExaminationService(money float64, openId string) (int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	examination, err := dao.GetUserExamination(tx, userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if examination.UserRemainder < money {
		return constant.StatusCodeInputError, errors.New("余额不足")
	}

	effectRows, err := dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: examination.UserCheckCount,
		UserRemainder:  tools.FloatRound(examination.UserRemainder-money, 2),
		UserCardType:   examination.UserCardType,
		UpdateTime:     time.Now(),
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.CreateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if effectRows == 0 {
		tx.Rollback()
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.CreateUserExamination effect rows = 0")
		return constant.StatusCodeInputError, errors.New("操作太频繁,请稍后重试")
	}
	err = dao.CreateExpenseCalendar(tx, &model.GuardianExpenseCalendar{
		ID:         tools.GenId(),
		UserId:     userInfo.ID,
		Money:      money,
		Status:     constant.ExpenseStatusRefund,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.CreateExpenseCalendar error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	tx.Commit()
	return constant.StatusCodeSuccess, nil
}
