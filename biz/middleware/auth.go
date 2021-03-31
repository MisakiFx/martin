package middleware

import (
	"net/http"

	"github.com/MisakiFx/martin/pkg/constant"
	"github.com/MisakiFx/martin/pkg/tools"
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
