package model

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
}
