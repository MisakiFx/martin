package dao

import (
	"time"

	"github.com/MisakiFx/martin/martin/pkg/model"

	"github.com/jinzhu/gorm"

	"github.com/MisakiFx/martin/martin/pkg/connection/mysql"
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

func GetCheckResultByBookingId(bookingId int64) (*model.GuardianCheckResult, error) {
	query := mysql.GetMysqlClient()
	var result model.GuardianCheckResult
	err := query.Table(result.TableName()).Where("booking_id = ?").First(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
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

func ListCheck(userId int64, page, size int) (int64, []model.GuardianBookingInfo, error) {
	query := mysql.GetMysqlClient()
	var bookingInfo model.GuardianBookingInfo
	result := make([]model.GuardianBookingInfo, 0)
	query = query.Table(bookingInfo.TableName()).Where("user_id = ?", userId).Order("create_time")
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	if page != 0 && size != 0 {
		query = query.Offset((page - 1) * size).Limit(size)
	}
	err = query.Find(&result).Error
	if err != nil {
		return 0, nil, err
	}
	return count, result, nil
}
