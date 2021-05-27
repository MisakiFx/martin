package dependencies

import (
	"runtime/debug"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_05vVAQeIbZWcnhLC478niV8rYKEoAl6y-6KGh8ai_-fbmTpsiczdn1rSDF7uTV_-P5tFGG1k4KBvJdc133Tlf_ukI47UamC2YVdtf5awp3avbXQBiJs_WTjMSaWLlb39etrYtJ2KSY1A38RxKCCaAGAZRR"

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
