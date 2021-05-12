package service

import (
	"errors"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/dao"
	"github.com/MisakiFx/martin/martin/pkg/tools"
)

func CheckAdmin(openId string) (int, error) {
	user, err := dao.CheckUserAdmin(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.CheckAdmin->dao.CheckUserAdmin error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	if user == nil {
		tools.GetLogger().Debugf("service.CheckAdmin openId : %v is not admin or not exist", openId)
		return constant.StatusCodeAuthError, errors.New("账户存在或没有管理员权限")
	}
	return constant.StatusCodeSuccess, nil
}
