package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/MisakiFx/martin/martin/pkg/service"

	"github.com/MisakiFx/martin/martin/pkg/model"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/tools"
	"github.com/gin-gonic/gin"
)

func checkBooingCheckReq(req *model.BookingCheckReq) error {
	startTime, err := time.ParseInLocation(constant.TimeFormatString, req.StartTime, tools.LocGloble)
	if err != nil || (startTime.Hour() != 8 && startTime.Hour() != 10 && startTime.Hour() != 14 && startTime.Hour() != 16) || startTime.Minute() != 0 || startTime.Second() != 0 {
		return errors.New("时间格式不合要求")
	}
	threeDaysLater := time.Now().Add(time.Hour * 24 * 3)
	limitTime := time.Date(threeDaysLater.Year(), threeDaysLater.Month(), threeDaysLater.Day(), 23, 59, 59, 0, tools.LocGloble)
	if limitTime.Sub(startTime) < 0 {
		return errors.New("只能预约近三天内的体检时间")
	}
	if startTime.Sub(time.Now()) < 0 {
		return errors.New("预约的体检时间已过期")
	}
	for _, project := range req.CheckProject {
		if _, ok := model.CheckProjectMap[project]; !ok {
			return errors.New("体检的项目不存在")
		}
	}
	if req.PayType != constant.PayTypeRemainder && req.PayType != constant.PayTypeCheckCount {
		return errors.New("付款方式不正确")
	}
	return nil
}

func BookingCheck(c *gin.Context) {
	tools.GetLogger().Infof("handler.BookingCheck path : %v", c.Request.URL.String())
	var req model.BookingCheckReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		tools.GetLogger().Errorf("handler.BookingCheck bind json error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	err = checkBooingCheckReq(&req)
	if err != nil {
		tools.GetLogger().Errorf("handler.BookingCheck check req error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  err.Error(),
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
	bookingId, statusCode, err := service.BookingCheckService(&req, openId)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.GetExaminationInfo->service.BookingCheckService error : %v", err)
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
			"id": bookingId,
		},
	})
}

func ListCheck(c *gin.Context) {
	tools.GetLogger().Infof("handler.ListCheck path : %v", c.Request.URL.String())
	page, size, err := getPageSizeFromQuery(c)
	if err != nil {
		tools.GetLogger().Errorf("handler.ListCheck get page size error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
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
	count, list, statusCode, err := service.ListCheckService(openId, page, size)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.ListCheck->service.ListCheckService error : %v", err)
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
			"list":  list,
			"count": count,
		},
	})
}

func GetCheckResult(c *gin.Context) {
	tools.GetLogger().Infof("handler.ListCheck path : %v", c.Request.URL.String())
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
	bookingIdString := c.Param("id")
	bookingId, err := strconv.ParseInt(bookingIdString, 10, 64)
	if err != nil {
		tools.GetLogger().Errorf("handler.GetCheckResult parse id error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	result, statusCode, err := service.GetCheckResultService(openId, bookingId)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.GetCheckResult->service.GetCheckResultService error : %v", err)
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
			"result": result,
		},
	})
}
