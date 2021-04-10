package tools

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/MisakiFx/martin/pkg/constant"

	"github.com/MisakiFx/martin/pkg/dependencies"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

var AccessToken string = "42_zW_EXSO_cKCKP6dYBkTRiDj8JnndxCyDbsRgjmuoIaB_8n2KKiI8VrlpYoW3aQUp4_JG_-o3STGr8AnUqOA1xiyqgXQ2cFnEopto4NTyFwsi8wkcboMH-uDZaCKsEBO-G03ISlfHo5PetU_0FSNeAIAMZU"
var JsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
var sugarLogger *zap.SugaredLogger
var LocGloble *time.Location

func getAccessTokenInit() {
	ticker := time.NewTicker(time.Minute * 110)
	accessToken, err := dependencies.GetAccessToken()
	if err != nil {
		GetLogger().Infof("Init accessToken error : %v", err)
		debug.PrintStack()
		panic(err)
	}
	AccessToken = accessToken
	for {
		<-ticker.C
		accessToken, err := dependencies.GetAccessToken()
		if err != nil {
			GetLogger().Errorf("Init accessToken error : %v", err)
			continue
		}
		AccessToken = accessToken
	}
}

func locationInit() {
	loc, err := time.LoadLocation(constant.TimeLocation)
	if err != nil {
		panic(fmt.Sprintf("Init time error : %v", err))
	}
	LocGloble = loc
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
	locationInit()
}
