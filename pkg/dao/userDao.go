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

func GetUserInfo(openId string) (*model.GuardianUserInfo, error) {
	query := mysql.GetMysqlClient()
	var user model.GuardianUserInfo
	err := query.Where("open_id = ?", openId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserBaseInfo(user *model.GuardianUserInfo) error {
	query := mysql.GetMysqlClient()
	err := query.Table(user.TableName()).Where("open_id = ?", user.OpenId).Updates(map[string]interface{}{
		"user_name":    user.UserName,
		"phone_number": user.PhoneNumber,
		"user_gender":  user.UserGender,
		"update_time":  user.UpdateTime,
	}).Error
	return err
}
