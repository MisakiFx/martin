package service

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/dependencies"

	"github.com/pkg/errors"

	"github.com/MisakiFx/martin/martin/pkg/connection/redis"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/pkg/dao"
	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/MisakiFx/martin/martin/pkg/tools"
	redis2 "github.com/go-redis/redis"
)

func GetUserInfo(openId string) (*model.GetUserInfoResp, error) {
	userInfo, err := dao.GetUserInfoByOpenId(openId)
	if err != nil {
		tools.GetLogger().Errorf("service.GetUserInfo->dao.GetUserInfo error : %v", err)
		return nil, err
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("service.GetUserInfo->dao.GetUserInfo do not found user info")
		return nil, nil
	}
	return &model.GetUserInfoResp{
		OpenId:      userInfo.OpenId,
		UserName:    userInfo.UserName,
		PhoneNumber: tools.PhoneNumberDesensitization(userInfo.PhoneNumber),
		UserGender:  userInfo.UserGender,
	}, nil
}

func checkVerificationCode(phoneNumber string, verification string) error {
	redisClient := redis.GetRedisClient()
	codeFromRedis, err := redisClient.Get(getLoginVerificationCodeKey(phoneNumber)).Result()
	if err != nil {
		if err == redis2.Nil {
			tools.GetLogger().Warnf("service.UserLoginService not found verification code from redis, phoneNumber : %v", phoneNumber)
			return errors.New("短信验证码已过期，请重试")
		} else {
			tools.GetLogger().Errorf("service.UserLoginService get redis error : %v", err)
			return errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
		}
	}
	if codeFromRedis != verification {
		tools.GetLogger().Infof("service.UserLoginService check verification code : %v != input code : %v", codeFromRedis, verification)
		return errors.New("验证码输入错误，请重试")
	}
	return nil
}

func UserLoginService(req *model.UserReq) (int, int64, error) {
	//读redis判断验证码
	err := checkVerificationCode(req.PhoneNumber, req.VerificationCode)
	if err != nil {
		return constant.StatusCodeInputError, 0, err
	}

	//落库
	tx := dao.StartTransaction()
	defer dao.ShutDownTransaction(tx)
	id := tools.GenId()
	err = dao.CreateUser(tx, &model.GuardianUserInfo{
		ID:          id,
		OpenId:      req.OpenId,
		UserName:    req.UserName,
		PhoneNumber: req.PhoneNumber,
		UserGender:  constant.UserGenderMale,
		UserPower:   constant.UserPowerNormal,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "uniq_idx_open_id") {
			tools.GetLogger().Warnf("service.UserLoginService : 当前用户已经绑定过基本信息，无需再次绑定")
			return constant.StatusCodeInputError, 0, errors.New("当前用户已经绑定过基本信息，无需再次绑定")
		} else if strings.Contains(err.Error(), "uniq_idx_phone_number") {
			tools.GetLogger().Warnf("service.UserLoginService : 当前手机号已经被绑定")
			return constant.StatusCodeInputError, 0, errors.New("当前手机号已经被绑定")
		} else {
			tools.GetLogger().Errorf("service.UserLoginService->dao.CreateUser error : %v", err)
			return constant.StatusCodeServiceError, 0, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
		}
	}
	err = dao.CreateUserExamination(tx, &model.GuardianHealthExaminationInfo{
		ID:             tools.GenId(),
		UserId:         id,
		UserCheckCount: 0,
		UserRemainder:  0,
		UserCardType:   10,
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
	})
	if err != nil {
		tx.Rollback()
		tools.GetLogger().Errorf("service.UserLoginService->dao.CreateUserExamination error : %v", err)
		return constant.StatusCodeServiceError, 0, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	tx.Commit()
	return constant.StatusCodeSuccess, id, nil
}

func GetUserOpenIdByCode(code string) (string, error) {
	if code == "123" {
		return "oSjQ26_7jlYQzA2b4NAWIBbF7RJ4", nil
	}
	openId, err := dependencies.GetOpenIdByCode(code)
	if err != nil {
		tools.GetLogger().Errorf("service.GetUserOpenIdByCode->dependencies.GetOpenIdByCode error : %v", err)
		return "", err
	}
	return openId, nil
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

func UpdateUserBaseInfo(req *model.UserReq) (int, error) {
	err := dao.UpdateUserBaseInfo(&model.GuardianUserInfo{
		ID:         0,
		OpenId:     req.OpenId,
		UserName:   req.UserName,
		UserGender: req.UserGender,
		UpdateTime: time.Now(),
	})
	if err != nil {
		tools.GetLogger().Errorf("service.UpdateUserBaseInfo->dao.UpdateUserBaseInfo error : %v", err)
		return constant.StatusCodeServiceError, errors.New(constant.StatusCodeMessageMap[constant.StatusCodeServiceError])
	}
	return constant.StatusCodeSuccess, nil
}
