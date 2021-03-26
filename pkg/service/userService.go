package service

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/MisakiFx/martin/pkg/dependencies"

	"github.com/pkg/errors"

	"github.com/MisakiFx/martin/pkg/connection/redis"

	"github.com/MisakiFx/martin/pkg/constant"

	"github.com/MisakiFx/martin/pkg/dao"
	"github.com/MisakiFx/martin/pkg/model"
	"github.com/MisakiFx/martin/pkg/tools"
	redis2 "github.com/go-redis/redis"
)

func UserLoginService(req *model.LoginReq) (int, int64, error) {

	//读redis判断验证码
	redisClient := redis.GetRedisClient()
	codeFromRedis, err := redisClient.Get(getLoginVerificationCodeKey(req.PhoneNumber)).Result()
	if err != nil {
		if err == redis2.Nil {
			tools.GetLogger().Warnf("service.UserLoginService not found verification code from redis, phoneNumber : %v", req.PhoneNumber)
			return constant.StatusCodeInputError, 0, errors.New("短信验证码已过期，请重试")
		} else {
			tools.GetLogger().Errorf("service.UserLoginService get redis error : %v", err)
			return constant.StatusCodeServiceError, 0, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
		}
	}
	if codeFromRedis != req.VerificationCode {
		tools.GetLogger().Infof("service.UserLoginService check verification code : %v != input code : %v", codeFromRedis, req.VerificationCode)
		return constant.StatusCodeInputError, 0, errors.New("验证码输入错误，请重试")
	}

	//落库
	id := tools.GenId()
	err = dao.CreateUser(&model.GuardianUserInfo{
		ID:          id,
		OpenId:      req.OpenId,
		UserName:    req.UserName,
		PhoneNumber: req.PhoneNumber,
		UserGender:  req.UserGender,
		UserPower:   req.UserPower,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		tools.GetLogger().Errorf("service.UserLoginService->dao.CreateUser error : %v", err)
		return constant.StatusCodeServiceError, 0, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}

	return constant.StatusCodeSuccess, id, nil
}

func getLoginVerificationCodeKey(phoneNumber string) string {
	return constant.VerificationCodeRedisKey + phoneNumber
}

func LoginVerificationCode(phoneNumber string) error {
	//生成验证码
	rand.Seed(time.Now().UnixNano())
	code := rand.Int63n(1000000)
	codeStr := strconv.FormatInt(code, 10)

	//发短信
	err := dependencies.SendMessage(phoneNumber, codeStr)
	if err != nil {
		return err
	}

	//写redis
	redisClient := redis.GetRedisClient()
	err = redisClient.Set(getLoginVerificationCodeKey(phoneNumber), codeStr, time.Minute*10).Err()
	if err != nil {
		tools.GetLogger().Errorf("service.LoginVerificationCode set redis error : %v", err)
		return err
	}
	return nil
}
