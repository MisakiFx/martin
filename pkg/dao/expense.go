package dao

import (
	"time"

	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/jinzhu/gorm"
)

func CreateExpenseCalendar(tx *gorm.DB, expense *model.GuardianExpenseCalendar) error {
	if tx == nil {
		tx = mysql.GetMysqlClient()
	}
	err := tx.Create(expense).Error
	return err
}

func ListExpenseCalendar(userId int64, page, size int, oldTime time.Time) (int64, []model.GuardianExpenseCalendar, error) {
	query := mysql.GetMysqlClient()
	var expenses []model.GuardianExpenseCalendar
	var expense model.GuardianExpenseCalendar
	query = query.Table(expense.TableName()).Where("user_id = ? AND create_time >= ?", userId, oldTime).Order("create_time DESC")
	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	if page != 0 && size != 0 {
		query = query.Offset((page - 1) * size).Limit(size)
	}
	err = query.Find(&expenses).Error
	if err != nil {
		return 0, nil, err
	}
	return count, expenses, nil
}
