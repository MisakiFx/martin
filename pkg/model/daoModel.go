package model

import "time"

type GuardianUserInfo struct {
	ID          int64     `gorm:"column:id" json:"id" form:"id"`
	OpenId      string    `gorm:"column:open_id" json:"open_id" form:"open_id"`
	UserName    string    `gorm:"column:user_name" json:"user_name" form:"user_name"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number" form:"phone_number"`
	UserGender  int       `gorm:"column:user_gender" json:"user_gender" form:"user_gender"`
	UserPower   int       `gorm:"column:user_power" json:"user_power" form:"user_power"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func (g *GuardianUserInfo) TableName() string {
	return "guardian_user_info"
}

type GuardianHealthExaminationInfo struct {
	ID             int64     `gorm:"column:id" json:"id" form:"id"`
	UserId         int64     `gorm:"column:user_id" json:"user_id" form:"user_id"`
	UserCheckCount int       `gorm:"column:user_check_count" json:"user_check_count" form:"user_check_count"`
	UserRemainder  float64   `gorm:"column:user_remainder" json:"user_remainder" form:"user_remainder"`
	UserCardType   int       `gorm:"column:user_card_type" json:"user_card_type" form:"user_card_type"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func (g *GuardianHealthExaminationInfo) TableName() string {
	return "guardian_health_examination_info"
}

type GuardianExpenseCalendar struct {
	ID         int64     `gorm:"column:id" json:"id" form:"id"`
	UserId     int64     `gorm:"column:user_id" json:"user_id" form:"user_id"`
	Money      float64   `gorm:"column:money" json:"money" form:"money"`
	Status     int       `gorm:"column:status" json:"status" form:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func (g *GuardianExpenseCalendar) TableName() string {
	return "guardian_expense_calendar"
}
