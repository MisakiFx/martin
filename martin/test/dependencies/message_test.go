package dependencies

import (
	"log"
	"testing"
	"time"

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
