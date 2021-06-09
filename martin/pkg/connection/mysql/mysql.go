package mysql

import (
	"time"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/pkg/dependencies"

	"github.com/MisakiFx/martin/martin/pkg/tools"

	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron/v3"
)

var mysqlClient *gorm.DB

const dsnHostIPString string = "root:1@tcp(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"

func InitCronJob() {
	c := cron.New()

	_, err := c.AddFunc("0-59/10 * * * *", func() {
		CronJobDo()
	})
	if err != nil {
		panic(err)
	}
	c.Start()
}

func Init() {
	db, err := gorm.Open("mysql", dsnHostIPString)
	if err != nil {
		panic(err)
	}
	mysqlClient = db
	InitCronJob()
}

func GetMysqlClient() *gorm.DB {
	return mysqlClient.Debug()
}

func CronJobDo() {
	query := GetMysqlClient()
	err := query.Table((&model.GuardianBookingInfo{}).TableName()).Where("end_time <= ?", time.Now()).Updates(map[string]interface{}{
		"status": constant.CheckEndStatus,
	}).Error
	if err != nil {
		tools.GetLogger().Errorf("更新已过期体检信息失败 : %v", err)
	}
	afterTenMinutesTime := time.Now().Add(time.Minute * 10)
	checkStarting := make([]model.GuardianBookingInfo, 0)
	err = query.Table((&model.GuardianBookingInfo{}).TableName()).Where("start_time = ?", afterTenMinutesTime).Find(&checkStarting).Error
	if err != nil {
		tools.GetLogger().Errorf("时间提醒服务失败 : %v", err)
	}
	for _, check := range checkStarting {
		tools.GetLogger().Debugf("userId : %v, check will start", check.UserId)
		var userInfo model.GuardianUserInfo
		err := query.Table((&model.GuardianUserInfo{}).TableName()).Where("id = ?", check.UserId).Find(&userInfo).Error
		if err != nil {
			tools.GetLogger().Errorf("CronJobDo get user info error : %v", err)
			continue
		}
		//调用发送模版消息
		err = dependencies.SendTemplateMessage(userInfo.OpenId, constant.TemplateIdCheckWillStart, "http://82.156.35.184:8080/me", map[string]string{
			"checkTime": afterTenMinutesTime.Format(constant.TimeFormatString2),
		})
		if err != nil {
			tools.GetLogger().Errorf("mysql.CronJobDo->dependencies.SendTemplateMessage error : %v", err)
		}
	}
}
