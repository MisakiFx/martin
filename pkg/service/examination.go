package service

import (
	"errors"
	"time"

	"github.com/MisakiFx/martin/pkg/constant"

	"github.com/MisakiFx/martin/pkg/dao"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/MisakiFx/martin/pkg/tools"
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
	examinationOld, err := dao.GetUserExamination(userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.GetUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}

	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	//更新用户体检卡信息
	err = dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: examinationOld.UserCheckCount + examination.CheckCount,
		UserRemainder:  examinationOld.UserRemainder + examination.Remainder,
		UserCardType:   examination.CardType,
		UpdateTime:     time.Now(),
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.BuyExaminationService->dao.CreateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
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
	examination, err := dao.GetUserExamination(userInfo.ID)
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

func RefundExaminationService(req *model.RefundExamination, openId string) (int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserInfoByOpenId error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.BuyExaminationService user not found")
		return constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	examination, err := dao.GetUserExamination(userInfo.ID)
	if err != nil {
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.GetUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if examination.UserRemainder < req.Money {
		return constant.StatusCodeInputError, errors.New("余额不足")
	}

	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	err = dao.UpdateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		UserId:         userInfo.ID,
		UserCheckCount: examination.UserCheckCount,
		UserRemainder:  examination.UserRemainder - req.Money,
		UserCardType:   examination.UserCardType,
		UpdateTime:     time.Now(),
	}, userInfo.ID)
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.RefundExaminationService->dao.CreateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	err = dao.CreateExpenseCalendar(tx, &model.GuardianExpenseCalendar{
		ID:         tools.GenId(),
		UserId:     userInfo.ID,
		Money:      req.Money,
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
