package middleware

import (
	"net/http"

	"github.com/MisakiFx/martin/martin/pkg/service"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/tools"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	openId := c.GetHeader("open_id")
	if openId == "" {
		tools.GetLogger().Errorf("middleware.Auth get user openId from header error")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	c.Set(constant.UserOpenIdContextKey, openId)
}

func AdminAuth(c *gin.Context) {
	openIdInterface, ok := c.Get(constant.UserOpenIdContextKey)
	openId, ok2 := openIdInterface.(string)
	if !ok || !ok2 || openId == "" {
		tools.GetLogger().Errorf("handler.BuyExamination get user info from context error")
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeAuthError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeAuthError],
		})
		return
	}
	code, err := service.CheckAdmin(openId)
	if code != constant.StatusCodeSuccess && err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}
}
