package model

import "time"

type UserReq struct {
	OpenId           string `json:"open_id"`
	UserName         string `json:"user_name"`
	PhoneNumber      string `json:"phone_number"`
	UserGender       int    `json:"user_gender"`
	UserPower        int    `json:"user_power"`
	VerificationCode string `json:"verification_code"`
}

type GetUserInfoResp struct {
	OpenId      string `json:"open_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	UserGender  int    `json:"user_gender"`
}

type BuyExaminationReq struct {
	ExaminationId int `json:"examination_id"`
}

type GetExaminationInfoResp struct {
	CheckCount int     `json:"check_count"`
	Remainder  float64 `json:"remainder"`
	CardType   int     `json:"card_type"`
}

type ListExpenseCalendarResp struct {
	Money      float64   `json:"money"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
}

type RefundExamination struct {
	Money float64 `json:"money"`
}
