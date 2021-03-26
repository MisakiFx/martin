package model

type LoginReq struct {
	OpenId           string `json:"open_id"`
	UserName         string `json:"user_name"`
	PhoneNumber      string `json:"phone_number"`
	UserGender       int    `json:"user_gender"`
	UserPower        int    `json:"user_power"`
	VerificationCode string `json:"verification_code"`
}
