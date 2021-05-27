package dao

import (
	"github.com/MisakiFx/martin/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/model"
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
		"user_name":   user.UserName,
		"user_gender": user.UserGender,
		"update_time": user.UpdateTime,
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

func CheckUserAdmin(openId string) (*model.GuardianUserInfo, error) {
	query := mysql.GetMysqlClient()
	var user model.GuardianUserInfo
	err := query.Table(user.TableName()).Where("open_id = ? AND user_power = ?", openId, constant.UserPowerAdmin).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserInfoByPhoneNumber(phoneNumber string) (*model.GuardianUserInfo, error) {
	query := mysql.GetMysqlClient()
	var user model.GuardianUserInfo
	err := query.Table(user.TableName()).Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUserPower(openId string, power int) error {
	query := mysql.GetMysqlClient()
	var user model.GuardianUserInfo
	err := query.Table(user.TableName()).Where("open_id = ?", openId).Updates(map[string]interface{}{
		"user_power": power,
	}).Error
	return err
}
