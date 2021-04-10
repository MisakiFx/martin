package dao

import (
	"time"

	"github.com/MisakiFx/martin/pkg/model"

	"github.com/jinzhu/gorm"

	"github.com/MisakiFx/martin/pkg/connection/mysql"
)

func GetStartedBookingCheck(tx *gorm.DB, userId int64) (*model.GuardianBookingInfo, error) {
	if tx != nil {
		tx = tx.Set("gorm:query_option", "FOR UPDATE")
	} else {
		tx = mysql.GetMysqlClient()
	}
	var info model.GuardianBookingInfo
	err := tx.Where("user_id = ? AND status < ?", userId, 8).First(&info).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &info, nil
}

func CountBookedCheckTime(tx *gorm.DB, startTime time.Time) (int, error) {
	if tx != nil {
		tx = tx.Set("gorm:query_option", "FOR UPDATE")
	} else {
		tx = mysql.GetMysqlClient()
	}
	var count int
	var bookingCheck model.GuardianBookingInfo
	err := tx.Table(bookingCheck.TableName()).Where("start_time = ?", startTime).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateBookingCheck(tx *gorm.DB, booking *model.GuardianBookingInfo) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(booking).Error
	return err
}

func CreateCheckResult(tx *gorm.DB, result *model.GuardianCheckResult) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(result).Error
	return err
}
