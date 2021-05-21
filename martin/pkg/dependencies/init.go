package dependencies

import (
	"runtime/debug"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_yIIZcqSkjEUbFdvaYdYvvH9a3jnYIewV8et3eqv2QDTbTN40pQmQyOMP1UeW7KOtalJW5Y4p6DOHnSwF_x1NIZjrB5E_HlOi5q2ykT0jVB_g0EdW9UWny5yG8CDWPUEe6J_GK8If2TRBPGRjZAYhADATUS"

func getAccessTokenInit() {
	//ticker := time.NewTicker(time.Minute * 110)
	accessToken, err := GetAccessToken()
	if err != nil {
		tools.GetLogger().Infof("Init accessToken error : %v", err)
		debug.PrintStack()
		panic(err)
	}
	AccessToken = accessToken
	tools.GetLogger().Infof("Init get accessToken success, accessToken : %v", AccessToken)
	//for {
	//	<-ticker.C
	//	accessToken, err := GetAccessToken()
	//	if err != nil {
	//		tools.GetLogger().Errorf("Init accessToken error : %v", err)
	//		continue
	//	}
	//	AccessToken = accessToken
	//}
}

func Init() {
	getAccessTokenInit()
}
