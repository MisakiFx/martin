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
			err = dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdAdmin, "http://10.227.31.2:8080/#/admin/top/?code=123", map[string]string{})
			if err != nil {
				tools.GetLogger().Errorf("handler.ServicePost->dependencies.SendTemplateMessage error : %v", err)
			}
		case "管理员密钥":
			err := dao.UpdateUserPower(input["FromUserName"], constant.UserPowerAdmin)
			err = dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdBeAdmin, "http://10.227.31.2:8080/#/admin/top/?code=123", map[string]string{
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
			err := dependencies.SendTemplateMessage(input["FromUserName"], constant.TemplateIdSubscribe, "http://10.227.31.2:8080/#/introduce/?code=123", map[string]string{
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
