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
