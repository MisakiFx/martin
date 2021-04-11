package dao

import (
	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/jinzhu/gorm"
)

func CreateUser(tx *gorm.DB, user *model.GuardianUserInfo) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(user).Error
	return err
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

func GetUserInfoByOpenId(openId string) (*model.GuardianUserInfo, error) {
	query := mysql.GetMysqlClient()
	var user model.GuardianUserInfo
	err := query.Table(user.TableName()).Where("open_id = ?", openId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
