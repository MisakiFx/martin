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
		ID:             tools.GenId(),
		UserId:         userInfo.ID,
		UserCheckCount: examinationOld.UserCheckCount + examination.CheckCount,
		UserRemainder:  examinationOld.UserRemainder + examination.Remainder,
		UserCardType:   examination.CardType,
		CreateTime:     time.Now(),
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
