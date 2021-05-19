package dao

import (
	"fmt"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/model"

	"github.com/jinzhu/gorm"

	"github.com/MisakiFx/martin/martin/pkg/connection/mysql"
)

func GetUserBookingInfo(userId int64, bookingId int64) (*model.GuardianBookingInfo, error) {
	query := mysql.GetMysqlClient()
	var result model.GuardianBookingInfo
	err := query.Table(result.TableName()).Where("id = ? AND user_id = ?", bookingId, userId).Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
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

func GetCheckResultByBookingId(bookingId, userId int64) (*model.GuardianCheckResult, error) {
	query := mysql.GetMysqlClient()
	var result model.GuardianCheckResult
	err := query.Table(result.TableName()).Where("booking_id = ? AND user_id = ?", bookingId, userId).First(&result).Error
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
	query = query.Table(bookingInfo.TableName()).Where("user_id = ?", userId).Order("create_time DESC")
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

func deleteCheckBooking(tx *gorm.DB, userId, bookingId int64) (int64, error) {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	var result model.GuardianBookingInfo
	tx = tx.Delete(result, "id = ? AND user_id = ? AND status = 0", bookingId, userId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func deleteCheckResult(tx *gorm.DB, bookingId, userId int64) (int64, error) {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	var result model.GuardianCheckResult
	tx = tx.Delete(result, "booking_id = ? AND user_id = ?", bookingId, userId)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
func DeleteCheck(tx *gorm.DB, userId, bookingId int64) (int64, error) {
	rows, err := deleteCheckBooking(tx, userId, bookingId)
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, nil
	}
	rows, err = deleteCheckResult(tx, bookingId, userId)
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, nil
	}
	return rows, nil
}

func GetCheckInfo(bookingId, userId int64) (*model.GuardianBookingInfo, error) {
	query := mysql.GetMysqlClient()
	var res model.GuardianBookingInfo
	err := query.Table(res.TableName()).Where("id = ? AND user_id = ?", bookingId, userId).Find(&res).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func UpdateCheckStatus(bookingId int64, status int) error {
	query := mysql.GetMysqlClient()
	err := query.Table((&model.GuardianBookingInfo{}).TableName()).Where("id = ?", bookingId).Updates(map[string]interface{}{
		"status": status,
	}).Error
	return err
}

func GetLastCheckedProjectBooking(userId int64, project int) (*model.GuardianBookingInfo, error) {
	query := mysql.GetMysqlClient()
	var bookingInfo model.GuardianBookingInfo
	err := query.Table((&model.GuardianBookingInfo{}).TableName()).Where("user_id = ? AND status > ? AND check_project LIKE ?", userId, project, fmt.Sprintf("%%%v%%", project)).Order("start_time DESC").Limit(1).First(&bookingInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &bookingInfo, nil
}

func UpdateCheckResult(bookingId int64, result *model.GuardianCheckResult) error {
	query := mysql.GetMysqlClient()
	var res model.GuardianCheckResult
	err := query.Table(res.TableName()).Where("booking_id = ?", bookingId).Updates(map[string]interface{}{
		"internal":       result.Internal,
		"surgery":        result.Surgery,
		"ENT":            result.Ent,
		"SGPT":           result.Sgpt,
		"blood_glucode":  result.BloodGlucode,
		"blood_fat":      result.BloodFat,
		"renal_function": result.RenalFunction,
	}).Error
	return err
}
