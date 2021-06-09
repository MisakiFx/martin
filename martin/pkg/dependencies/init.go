package dependencies

import (
	"runtime/debug"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_hQ9b2c7VxCunka4XvSKHWh4HWGWJ3ZGJf01ARansSOcxf0_4JOF8996p1M09IcxfURqynYAZX3v8N-oRhznZdZrFBdgYrT7_do2OWqM7S-c-K9fy4Sm7NGW9W28pe_MqFHwHBAkX1IjQ98rXNLPeAEAFDH"

func getAccessTokenInit() {
	ticker := time.NewTicker(time.Minute * 110)
	accessToken, err := GetAccessToken()
	if err != nil {
		tools.GetLogger().Infof("Init accessToken error : %v", err)
		debug.PrintStack()
		panic(err)
	}
	AccessToken = accessToken
	tools.GetLogger().Infof("Init get accessToken success, accessToken : %v", AccessToken)
	go func() {
		for {
			<-ticker.C
			accessToken, err := GetAccessToken()
			if err != nil {
				tools.GetLogger().Errorf("Init accessToken error : %v", err)
				continue
			}
			AccessToken = accessToken
		}
	}()
}

func Init() {
	getAccessTokenInit()
}
