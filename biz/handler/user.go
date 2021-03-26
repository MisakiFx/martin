package handler

import (
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/MisakiFx/martin/pkg/tools"

	"github.com/MisakiFx/martin/pkg/service"

	"github.com/MisakiFx/martin/pkg/constant"
	"github.com/MisakiFx/martin/pkg/model"

	"github.com/gin-gonic/gin"
)

func checkLoginReq(req *model.LoginReq) error {
	if req == nil {
		return errors.New("入参错误")
	}
	if len(req.OpenId) <= 0 {
		return errors.New("用户openId不能为空")
	}

	if len(req.UserName) <= 0 {
		return errors.New("用户名不能为空")
	}
	if utf8.RuneCountInString(req.UserName) > 255 {
		return errors.New("用户名超过255个字符")
	}

	if len(req.PhoneNumber) <= 0 {
		return errors.New("用户电话不能为空")
	}
	phoneNumberCount := utf8.RuneCountInString(req.PhoneNumber)
	if phoneNumberCount != 11 && phoneNumberCount != 8 {
		return errors.New("用户电话不合法")
	}

	if req.UserGender == 0 {
		return errors.New("用户性别不能为空")
	}
	if req.UserGender != constant.UserGenderMale && req.UserGender != constant.UserGenderFemale {
		return errors.New("用户性别不合法")
	}

	if req.UserPower == 0 {
		return errors.New("用户权限不能为空")
	}
	if req.UserPower != constant.UserPowerNormal && req.UserPower != constant.UserPowerAdmin {
		return errors.New("用户权限不合法")
	}

	if req.VerificationCode == "" {
		return errors.New("请输入验证码")
	}
	return nil
}

func UserLogin(c *gin.Context) {
	tools.GetLogger().Infof("handler.UserLogin url : %v", c.Request.URL.String())

	var loginReq model.LoginReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		tools.GetLogger().Errorf("handler.UserLogin bind json error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	err = checkLoginReq(&loginReq)
	if err != nil {
		tools.GetLogger().Errorf("handler.UserLogin check input error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  err,
		})
		return
	}
	statusCode, id, err := service.UserLoginService(&loginReq)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.UserLogin->service.UserLoginService error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": statusCode,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.StatusCodeSuccess,
		"msg":  constant.StatusCodeMessageMap[constant.StatusCodeSuccess],
		"data": gin.H{
			"id": id,
		},
	})
}

func VerificationCode(c *gin.Context) {
	tools.GetLogger().Infof("handler.VerificationCode url : %v", c.Request.URL.String())
	phoneNumber := c.DefaultQuery("phone", "")
	if phoneNumber == "" {
		tools.GetLogger().Errorf("handler.VerificationCode phone is empty : %v", phoneNumber)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	err := service.LoginVerificationCode(phoneNumber)
	if err != nil {
		tools.GetLogger().Errorf("handler.UserLogin->service.UserLoginService error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeServiceError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeServiceError],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.StatusCodeSuccess,
		"msg":  constant.StatusCodeMessageMap[constant.StatusCodeSuccess],
	})
}
