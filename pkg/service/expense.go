package service

import (
	"errors"
	"time"

	"github.com/MisakiFx/martin/pkg/constant"
	"github.com/MisakiFx/martin/pkg/dao"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/MisakiFx/martin/pkg/tools"
)

func ListExpenseCalendarService(openId string, page int, size int) ([]model.ListExpenseCalendarResp, int64, int, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.ListExpenseCalendarService->dao.GetUserInfoByOpenId error : %v", err)
		return nil, 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.ListExpenseCalendarService user not found")
		return nil, 0, constant.StatusCodeAuthError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeAuthError])
	}
	//操作记录只取过去30天的
	oldTime := time.Now().AddDate(0, 0, -30)
	count, expenses, err := dao.ListExpenseCalendar(userInfo.ID, page, size, oldTime)
	if err != nil {
		tools.GetLogger().Errorf("service.ListExpenseCalendarService->dao.ListExpenseCalendar error : %v", err)
		return nil, 0, constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	list := make([]model.ListExpenseCalendarResp, 0)
	for _, expense := range expenses {
		list = append(list, model.ListExpenseCalendarResp{
			Money:      expense.Money,
			Status:     expense.Status,
			CreateTime: expense.CreateTime.Format(constant.TimeFormatString),
		})
	}
	return list, count, constant.StatusCodeSuccess, nil
}
