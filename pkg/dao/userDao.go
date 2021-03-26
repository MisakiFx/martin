package dao

import (
	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/model"
)

func CreateUser(user *model.GuardianUserInfo) error {
	query := mysql.GetMysqlClient()
	err := query.Create(user).Error
	return err
}
