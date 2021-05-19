package model

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
	CardType   float64 `json:"card_type"`
}

type ListExpenseCalendarResp struct {
	Money           float64 `json:"money"`
	Status          int     `json:"status"`
	ExaminationName string  `json:"examination_name"`
	CreateTime      string  `json:"create_time"`
}

type RefundExamination struct {
	Money string `json:"money"`
}
type BookingCheckReq struct {
	CheckProject []int  `json:"check_project"`
	StartTime    string `json:"start_time"`
	PayType      int    `json:"pay_type"`
}

type ListCheckResp struct {
	Id            int64   `json:"id"`
	CheckProject  []int   `json:"check_project"`
	StartTime     string  `json:"start_time"`
	EndTime       string  `json:"end_time"`
	Status        int     `json:"status"`
	CreateTime    string  `json:"create_time"`
	PayReminder   float64 `json:"pay_reminder"`
	PayCheckCount int     `json:"pay_check_count"`
}

type GetCheckResultResp struct {
	BookingId     int64  `json:"booking_id"`
	Projects      []int  `json:"projects"`
	Internal      string `json:"internal"`
	Surgery       string `json:"surgery"`
	ENT           string `json:"ent"`
	SGPT          string `json:"sgpt"`
	BloodGlucode  string `json:"blood_glucode"`
	BloodFat      string `json:"blood_fat"`
	RenalFunction string `json:"renal_function"`
}

type CancelCheckBookingReq struct {
	BookingId int64 `json:"booking_id"`
}

type CheckStartReq struct {
	PhoneNumber string `json:"phone_number"`
}

type CheckFinishReq struct {
	PhoneNumber   string `json:"phone_number"`
	FinishProject int    `json:"finish_project"`
}

type CheckResultReq struct {
	PhoneNumber  string `json:"phone_number"`
	CheckProject int    `json:"check_project"`
	CheckResult  string `json:"check_result"`
}
