package dependencies

import (
	"runtime/debug"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_fICR1ZoVj_Y_pxPteb2PSrekI6LyAlEktXgrIs3lwUFAnHKP30AAtULVygVHO8ID_b3fgaYrqhl3cMukXIeuZa1ED8SlOUqfKxTueleHrY7p-428pUeyRw6KWgFBAqYSFE_5lygmm7Gka6K4KJVbAAAGFR"

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
