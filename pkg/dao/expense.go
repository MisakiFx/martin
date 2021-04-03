package dao

import (
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
