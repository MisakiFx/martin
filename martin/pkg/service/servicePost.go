package service

import (
	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/dao"
	"github.com/MisakiFx/martin/martin/pkg/dependencies"
	"github.com/MisakiFx/martin/martin/pkg/tools"
)

func ServicePost(input map[string]string) error {
	tools.GetLogger().Infof("input : %+v", input)
	switch input["MsgType"] {
	case "text":
		switch input["Content"] {
		case "管理员":
			user, err := dao.CheckUserAdmin(input["FromUserName"])
			if err != nil {
				tools.GetLogger().Errorf("handler.ServicePost->dao.CheckUserAdmin error : %v", err)
			}
			if user == nil {
				tools.GetLogger().Warnf("handler.ServicePost user is not admin")
				return nil
			}
			err = dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdAdmin, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2Fadmin%2Ftop&response_type=code&scope=snsapi_base&state=1#wechat_redirect", map[string]string{})
			if err != nil {
				tools.GetLogger().Errorf("handler.ServicePost->dependencies.SendTemplateMessage error : %v", err)
			}
		case "管理员密钥":
			err := dao.UpdateUserPower(input["FromUserName"], constant.UserPowerAdmin)
			err = dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdBeAdmin, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2Fadmin%2Ftop&response_type=code&scope=snsapi_base&state=1#wechat_redirect", map[string]string{
				"word": "管理员",
			})
			if err != nil {
				tools.GetLogger().Errorf("handler.ServicePost->dependencies.SendTemplateMessage error : %v", err)
			}
		default:
		}
	case "event":
		switch input["Event"] {
		case "subscribe":
			err := dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdSubscribe, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx33cc6387acefe650&redirect_uri=http%3A%2F%2F82.156.35.184%3A8080%2Fintroduce&response_type=code&scope=snsapi_base&state=1#wechat_redirect", map[string]string{
				"introduce": "使用说明",
			})
			if err != nil {
				tools.GetLogger().Errorf("handler.ServicePost->dependencies.SendTemplateMessage error : %v", err)
			}
		default:
		}
	default:
		tools.GetLogger().Infof("default")
	}
	return nil
}
