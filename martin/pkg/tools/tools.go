package tools

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"go.uber.org/zap"

	"github.com/MisakiFx/martin/martin/pkg/constant"
)

func CheckToken(signature, timestamp, nonce string) bool {
	args := []string{
		constant.Token,
		nonce,
		timestamp,
	}
	sort.Strings(args)
	stringNotSha1 := ""
	for _, arg := range args {
		stringNotSha1 += arg
	}
	h := sha1.New()
	h.Write([]byte(stringNotSha1))
	stringSha1 := fmt.Sprintf("%x", h.Sum(nil))
	if stringSha1 != signature {
		return false
	}
	return true
}

func GenId() int64 {
	timeNow := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	randMath := rand.Int63n(1000000)
	return timeNow*1000000 + randMath
}

func GetLogger() *zap.SugaredLogger {
	return sugarLogger
}

func PhoneNumberDesensitization(phoneNumber string) string {
	prefix := phoneNumber[:3]
	suffix := phoneNumber[7:]
	return prefix + "xxxx" + suffix
}
