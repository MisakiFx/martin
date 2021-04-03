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
