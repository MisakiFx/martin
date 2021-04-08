package handler

import (
	"net/http"

	"github.com/MisakiFx/martin/pkg/service"

	"github.com/MisakiFx/martin/pkg/constant"
	"github.com/MisakiFx/martin/pkg/tools"
	"github.com/gin-gonic/gin"
)

func ListExpenseCalendar(c *gin.Context) {
	tools.GetLogger().Infof("handler.ListExpenseCalendar path : %v", c.Request.URL.String())
	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	page, size, err := getPageSizeFromQuery(c)
	if err != nil {
		tools.GetLogger().Errorf("handler.ListExpenseCalendar get page size error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.ListExpenseCalendar get user info from context error")
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	list, count, statusCode, err := service.ListExpenseCalendarService(openId, page, size)
	if statusCode != constant.StatusCodeSuccess {
		tools.GetLogger().Errorf("handler.ListExpenseCalendarService->service.ListExpenseCalendarService error : %v", err)
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
