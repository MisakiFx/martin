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
	UserCardType   float64   `gorm:"column:user_card_type" json:"user_card_type" form:"user_card_type"`
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

type GuardianBookingInfo struct {
	ID           int64     `gorm:"column:id" json:"id" form:"id"`
	UserId       int64     `gorm:"column:user_id" json:"user_id" form:"user_id"`
	CheckProject string    `gorm:"column:check_project" json:"check_project" form:"check_project"`
	StartTime    time.Time `gorm:"column:start_time" json:"start_time" form:"start_time"`
	EndTime      time.Time `gorm:"column:end_time" json:"end_time" form:"end_time"`
	Pay          string    `gorm:"column:pay" json:"pay" form:"pay"`
	Status       int       `gorm:"column:status" json:"status" form:"status"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func (g *GuardianBookingInfo) TableName() string {
	return "guardian_booking_info"
}

type GuardianCheckResult struct {
	ID            int64     `gorm:"column:id" json:"id" form:"id"`
	BookingId     int64     `gorm:"column:booking_id" json:"booking_id" form:"booking_id"`
	Internal      string    `gorm:"column:internal" json:"internal" form:"internal"`
	Surgery       string    `gorm:"column:surgery" json:"surgery" form:"surgery"`
	Ent           string    `gorm:"column:ent" json:"ent" form:"ent"`
	Sgpt          string    `gorm:"column:sgpt" json:"sgpt" form:"sgpt"`
	BloodGlucode  string    `gorm:"column:blood_glucode" json:"blood_glucode" form:"blood_glucode"`
	BloodFat      string    `gorm:"column:blood_fat" json:"blood_fat" form:"blood_fat"`
	RenalFunction string    `gorm:"column:renal_function" json:"renal_function" form:"renal_function"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time" form:"update_time"`
}

func (g *GuardianCheckResult) TableName() string {
	return "guardian_check_result"
}
