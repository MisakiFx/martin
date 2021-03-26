package tools

import (
	"os"
	"runtime/debug"
	"time"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

var AccessToken string = "42_zW_EXSO_cKCKP6dYBkTRiDj8JnndxCyDbsRgjmuoIaB_8n2KKiI8VrlpYoW3aQUp4_JG_-o3STGr8AnUqOA1xiyqgXQ2cFnEopto4NTyFwsi8wkcboMH-uDZaCKsEBO-G03ISlfHo5PetU_0FSNeAIAMZU"
var JsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
var sugarLogger *zap.SugaredLogger

func getAccessTokenInit() {
	ticker := time.NewTicker(time.Minute * 110)
	accessToken, err := GetAccessToken()
	if err != nil {
		GetLogger().Infof("Init accessToken error : %v", err)
		debug.PrintStack()
		panic(err)
	}
	AccessToken = accessToken
	for {
		<-ticker.C
		accessToken, err := GetAccessToken()
		if err != nil {
			GetLogger().Errorf("Init accessToken error : %v", err)
			continue
		}
		AccessToken = accessToken
	}
}

func loggerInit() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}
func Init() {
	//getAccessTokenInit()
	loggerInit()
}
