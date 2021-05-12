package tools

import (
	"fmt"
	"os"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"

	jsoniter "github.com/json-iterator/go"
)

var JsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
var sugarLogger *zap.SugaredLogger
var LocGloble *time.Location

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
	loggerInit()
	locationInit()
}
