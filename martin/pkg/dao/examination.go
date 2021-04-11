package dao

import (
	"time"

	"github.com/MisakiFx/martin/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/jinzhu/gorm"
)

func CreateUserExamination(tx *gorm.DB, examination *model.GuardianHealthExaminationInfo) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(examination).Error
	return err
}

func UpdateUserExamination(tx *gorm.DB, examination *model.GuardianHealthExaminationInfo, userId int64) (int64, error) {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	tx = tx.Table(examination.TableName()).Where("user_id = ? AND update_time <= ?", userId, examination.UpdateTime).Updates(map[string]interface{}{
		"user_check_count": examination.UserCheckCount,
		"user_remainder":   examination.UserRemainder,
		"user_card_type":   examination.UserCardType,
		"update_time":      time.Now(),
	})
	return tx.RowsAffected, tx.Error
}

func GetUserExamination(tx *gorm.DB, userId int64) (*model.GuardianHealthExaminationInfo, error) {
	if tx != nil {
		tx = tx.Set("gorm:query_option", "FOR UPDATE")
	} else {
		tx = mysql.GetMysqlClient()
	}
	var examination model.GuardianHealthExaminationInfo
	err := tx.Table(examination.TableName()).Where("user_id = ?", userId).First(&examination).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &examination, nil
}
