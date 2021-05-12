package dependencies

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MisakiFx/martin/martin/pkg/tools"

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

func SendTemplateMessage(openId string, templateId string, templateValue map[string]string) error {
	bodyMap := make(map[string]interface{}, 0)
	bodyMap["touser"] = openId
	bodyMap["template_id"] = templateId
	bodyMap["topcolor"] = "#FF0000"
	dataMap := make(map[string]map[string]string, 0)
	for field, value := range templateValue {
		dataMap[field] = map[string]string{
			"value": value,
			"color": "#173177",
		}
	}
	bodyMap["data"] = dataMap
	bodyString, err := tools.JsonIter.Marshal(bodyMap)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(bodyString))
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%v", AccessToken), "application/json", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("http resp code != 200")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respBodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(respBody, &respBodyMap)
	if err != nil {
		return err
	}
	if respBodyMap["errcode"].(float64) != 0 {
		return errors.New(respBodyMap["errmsg"].(string))
	}
	return nil
}
