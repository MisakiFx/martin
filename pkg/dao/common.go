package dao

import (
	"github.com/MisakiFx/martin/pkg/connection/mysql"
	"github.com/MisakiFx/martin/pkg/tools"

	"github.com/jinzhu/gorm"
)

func StartTransaction() *gorm.DB {
	db := mysql.GetMysqlClient()
	tx := db.Begin()
	return tx
}

func ShutDownTransaction(tx *gorm.DB) {
	if err := recover(); err != nil && tx != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("tx rollback because of %s", err)
	}
}
