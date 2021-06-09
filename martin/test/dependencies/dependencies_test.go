package dependencies

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/tools"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/pkg/dependencies"
)

func TestSendMessage(t *testing.T) {
	err := dependencies.SendMessage("18192243121", "6666")
	if err != nil {
		panic(err)
	}
	log.Printf("成功")
}

func TestGetAccessToken(t *testing.T) {
	accessToken, err := dependencies.GetAccessToken()
	if err != nil {
		log.Fatalf("tools.GetAccessToken error : %v", err)
		return
	}
	log.Printf("accessToken : %v", accessToken)
	time.Sleep(time.Second)
}

func TestSendTemplateMessage(t *testing.T) {
	err := dependencies.SendTemplateMessage("oSjQ26_7jlYQzA2b4NAWIBbF7RJ4", constant.TemplateIdCheckWillStart, "http://10.227.31.2:8080/#/me?code=123", map[string]string{
		"checkTime": time.Now().Format(constant.TimeFormatString2),
	})
	if err != nil {
		tools.GetLogger().Errorf("error : %v", err)
	}
	time.Sleep(time.Second)
}

func TestBuildMenu(t *testing.T) {
	body := `
{
    "button":[
        {
            "type":"view",
            "name":"使用介绍",
            "url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2F%23%2Fintroduce%2F&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
        },
        {
            "type":"view",
            "name":"体检商城",
            "url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2F%23%2Fhome%2F&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
        },
        {
            "name":"个人中心",
            "sub_button":[
                {
                    "type":"view",
                    "name":"个人信息",
                    "url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2F%23%2Fme%2F&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
                },
                {
                    "type":"view",
                    "name":"体检预约",
                    "url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F10.227.31.2%3A8080%2F%23%2Fcheck_booking%2F%3Fcode%3D123&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
                }
            ]
        }
    ]
}
	`
	resp, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%v", dependencies.AccessToken), "application/json", strings.NewReader(body))
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Printf("resp : %v", string(respBody))
}

func TestGetMenu(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/get?access_token=%v", dependencies.AccessToken))
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	fmt.Printf("resp : %v", string(respBody))
}
