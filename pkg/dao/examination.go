package dao

import (
	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/jinzhu/gorm"
)

func CreateUserExamination(tx *gorm.DB, examination *model.GuardianHealthExaminationInfo) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(examination).Error
	return err
}

func UpdateUserExamination(tx *gorm.DB, examination *model.GuardianHealthExaminationInfo, userId int64) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Where("user_id = ?", userId).Updates(map[string]interface{}{
		"user_check_count": examination.UserCheckCount,
		"user_remainder":   examination.UserRemainder,
		"user_card_type":   examination.UserCardType,
		"update_time":      examination.UpdateTime,
	}).Error
	return err
}

func GetUserExamination(userId int64) (*model.GuardianHealthExaminationInfo, error) {
	query := mysql.GetMysqlClient()
	var examination model.GuardianHealthExaminationInfo
	err := query.Table(examination.TableName()).Where("user_id = ?", userId).First(&examination).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &examination, nil
}
