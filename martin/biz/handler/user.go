package handler

import (
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/MisakiFx/martin/martin/pkg/tools"

	"github.com/MisakiFx/martin/martin/pkg/service"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/model"

	"github.com/gin-gonic/gin"
)

func checkUserReq(req *model.UserReq, isUpdate bool) error {
	if req == nil {
		return errors.New("入参错误")
	}
	if len(req.OpenId) <= 0 {
		return errors.New("无法识别用户身份，请从微信公众号打开")
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

	if req.VerificationCode == "" {
		return errors.New("请输入验证码")
	}
	return nil
}

func UserLogin(c *gin.Context) {
	tools.GetLogger().Infof("handler.UserLogin url : %v", c.Request.URL.String())

	var loginReq model.UserReq
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		tools.GetLogger().Errorf("handler.UserLogin bind json error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	err = checkUserReq(&loginReq, false)
	if err != nil {
		tools.GetLogger().Errorf("handler.UserLogin check input error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  err.Error(),
		})
		return
	}
	statusCode, id, err := service.UserLoginService(&loginReq)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.UserLogin->service.UserLoginService error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": statusCode,
			"msg":  err.Error(),
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

func GetUserInfo(c *gin.Context) {
	tools.GetLogger().Infof("handler.GetUserInfo url : %v", c.Request.URL.String())
	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.GetUserInfo get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	userInfo, err := service.GetUserInfo(openId)
	if err != nil {
		tools.GetLogger().Errorf("handler.GetUserInfo->service.GetUserInfo error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeServiceError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeServiceError],
		})
		return
	}
	if userInfo == nil {
		tools.GetLogger().Errorf("handler.GetUserInfo->service.GetUserInfo do not found user info")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.StatusCodeSuccess,
		"msg":  constant.StatusCodeMessageMap[constant.StatusCodeSuccess],
		"data": gin.H{
			"info": userInfo,
		},
	})
}

func GetUserOpenIdByCode(c *gin.Context) {
	tools.GetLogger().Infof("handler.GetUserOpenIdByCode url : %v", c.Request.URL.String())
	code := c.Param("code")
	if code == "" {
		tools.GetLogger().Errorf("handler.GetUserOpenIdByCode code is empty : %v", code)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	openId, err := service.GetUserOpenIdByCode(code)
	if err != nil {
		tools.GetLogger().Errorf("hendler.GetUserOpenIdByCode->service.GetUserOpenIdByCode error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeServiceError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeServiceError],
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.StatusCodeSuccess,
		"msg":  constant.StatusCodeMessageMap[constant.StatusCodeSuccess],
		"data": gin.H{
			"open_id": openId,
		},
	})
}

func UpdateUserInfo(c *gin.Context) {
	tools.GetLogger().Infof("handler.UpdateUserInfo url : %v", c.Request.URL.String())
	var req model.UserReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		tools.GetLogger().Errorf("handler.UpdateUserInfo bind model err : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.UpdateUserInfo get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	req.OpenId = openId
	err = checkUserReq(&req, true)
	if err != nil {
		tools.GetLogger().Errorf("handler.UpdateUserInfo check input error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  err.Error(),
		})
		return
	}
	statusCode, err := service.UpdateUserBaseInfo(&req)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.UpdateUserInfo->service.UpdateUserBaseInfo error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": statusCode,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": constant.StatusCodeSuccess,
		"msg":  constant.StatusCodeMessageMap[constant.StatusCodeSuccess],
	})
}
