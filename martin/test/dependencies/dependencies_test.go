package dependencies

import (
	"fmt"
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
	err := dependencies.SendTemplateMessage("oSjQ26_7jlYQzA2b4NAWIBbF7RJ4", constant.TemplateIdCheckWillStart, map[string]string{
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
          "name":"体检商城",
		  "url" : "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=https://1dce62b7b2b1.ngrok.io/#/home&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
      },
	{
          "type":"view",
          "name":"体检预约",
		  "url" : "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=https://1dce62b7b2b1.ngrok.io/#/home&response_type=code&scope=snsapi_base&state=1#wechat_redirect"
	}
	]}
	`
	http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%v", dependencies.AccessToken), "application/json", strings.NewReader(body))
}
