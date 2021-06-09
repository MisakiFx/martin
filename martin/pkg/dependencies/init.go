package dependencies

import (
	"runtime/debug"

	"github.com/MisakiFx/martin/martin/pkg/tools"
)

var AccessToken string = "45_r0jzXgiTz_5NESdLgZ9wXMK7243YFEIx-i1jKiIfxSLw9zsfTe79aJnt1PEd28i0xhqAlroUkIoFmzY_XJvNKpQlyA4kGlC-wz1sytpROCPP3m6qo2ziOrdAkt8pQuC1dpknO_fZnI4_ZHUjVKZdAEAYNC"

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
