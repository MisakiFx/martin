package middleware

import (
	"net/http"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/pkg/tools"
	"github.com/gin-gonic/gin"
)

func CheckToken(c *gin.Context) {
	signature := c.DefaultQuery("signature", "")
	timestamp := c.DefaultQuery("timestamp", "")
	nonce := c.DefaultQuery("nonce", "")
	isFixed := tools.CheckToken(signature, timestamp, nonce)
	if !isFixed {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeTokenCheckError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeTokenCheckError],
		})
		return
	}
}
