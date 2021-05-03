package handler

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/MisakiFx/martin/martin/pkg/service"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/model"
	"github.com/MisakiFx/martin/martin/pkg/tools"
	"github.com/gin-gonic/gin"
)

func BuyExamination(c *gin.Context) {
	tools.GetLogger().Infof("handler.BuyExamination path : %v", c.Request.URL.String())
	var req model.BuyExaminationReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		tools.GetLogger().Errorf("handler.BuyExamination bind json error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	if _, ok := model.ExaminationMap[req.ExaminationId]; !ok {
		tools.GetLogger().Errorf("handler.BuyExamination ExaminationId error, id : %v", req.ExaminationId)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  "商品类型不正确",
		})
		return
	}

	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.BuyExamination get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}

	statusCode, err := service.BuyExaminationService(&req, openId)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.BuyExamination->service.BuyExaminationService error : %v", err)
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

func GetExaminationInfo(c *gin.Context) {
	tools.GetLogger().Infof("handler.GetExaminationInfo path : %v", c.Request.URL.String())

	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.BuyExamination get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	examination, statusCode, err := service.GetExaminationInfoService(openId)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.GetExaminationInfo->service.GetExaminationInfoService error : %v", err)
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
			"info": examination,
		},
	})
}

func RefundExamination(c *gin.Context) {
	tools.GetLogger().Infof("handler.RefundExamination path : %v", c.Request.URL.String())
	var req model.RefundExamination
	err := c.ShouldBindJSON(&req)
	if err != nil {
		tools.GetLogger().Errorf("handler.RefundExamination bind json error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	moneyArray := strings.Split(req.Money, ".")
	if len(moneyArray) > 1 && utf8.RuneCountInString(moneyArray[1]) > 2 {
		tools.GetLogger().Warnf("handler.RefundExamination money float >= 2")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  "退款金额支持到小数点后两位",
		})
		return
	}
	moneyFloat, err := strconv.ParseFloat(req.Money, 64)
	if err != nil {
		tools.GetLogger().Warnf("handler.RefundExamination money parse error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	if moneyFloat <= 0 {
		tools.GetLogger().Warnf("handler.RefundExamination money can not <= 0")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  "退款金额不能小于等于0",
		})
		return
	}
	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.BuyExamination get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	statusCode, err := service.RefundExaminationService(moneyFloat, openId)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.RefundExamination->service.RefundExaminationService error : %v", err)
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
