package dependencies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/model"
)

func GetAccessToken() (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", constant.AppID, constant.Appsecret))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var accessTokenRes model.AccessToken
	err = json.Unmarshal(body, &accessTokenRes)
	if err != nil {
		return "", err
	}
	return accessTokenRes.AccessToken, nil
}

func GetOpenIdByCode(code string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%v&secret=%v&code=%v&grant_type=authorization_code", constant.AppID, constant.Appsecret, code))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var accessToken model.Oauth2AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}
	return accessToken.OpenId, nil
}
