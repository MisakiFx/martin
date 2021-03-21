package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlClient *gorm.DB

const dsnHostIPString string = "root:1@tcp(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"

func Init() {
	db, err := gorm.Open("mysql", dsnHostIPString)
	if err != nil {
		panic(err)
	}
	mysqlClient = db
}

func GetMysqlClient() *gorm.DB {
	return mysqlClient
}
