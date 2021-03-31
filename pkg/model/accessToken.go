package model

type AccessToken struct {
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
}

type Oauth2AccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}
