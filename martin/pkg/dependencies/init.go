package dependencies

import (
	"runtime/debug"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_hfcgevm6kobw5v6TWVSAksrfW0xLxjhhmXouFYzZKHaMHe6u2M18-YoQsA9VBaf4oCx4tjBIlYs_IzB-qWcC-bG-ET0PftvjlDePInVYI8s0z5TVLbXc8a2aPpqfeyz1vF5EK72aebDOoMxnXMPhAIAJYB"

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
