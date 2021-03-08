package tools

import (
	"log"
	"runtime/debug"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var AccessToken string = "42_zW_EXSO_cKCKP6dYBkTRiDj8JnndxCyDbsRgjmuoIaB_8n2KKiI8VrlpYoW3aQUp4_JG_-o3STGr8AnUqOA1xiyqgXQ2cFnEopto4NTyFwsi8wkcboMH-uDZaCKsEBO-G03ISlfHo5PetU_0FSNeAIAMZU"
var JsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

func Init() {
	ticker := time.NewTicker(time.Minute * 110)
	accessToken, err := GetAccessToken()
	if err != nil {
		log.Printf("Init accessToken error : %v", err)
		debug.PrintStack()
		panic(err)
	}
	AccessToken = accessToken
	for {
		<-ticker.C
		accessToken, err := GetAccessToken()
		if err != nil {
			log.Printf("Init accessToken error : %v", err)
			continue
		}
		AccessToken = accessToken
	}
}
