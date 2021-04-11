package handler

import (
	"net/http"

	"github.com/MisakiFx/martin/martin/pkg/tools"

	"github.com/MisakiFx/martin/martin/pkg/constant"
	"github.com/MisakiFx/martin/martin/pkg/model"

	"github.com/gin-gonic/gin"
)

//微信服务器发来的GET请求，只为了进行token验证
func ServiceGet(c *gin.Context) {
	tools.GetLogger().Infof("ServiceGet url : %v", c.Request.URL.String())
	echostr := c.DefaultQuery("echostr", "")
	c.String(http.StatusOK, "%s", echostr)
}

//微信服务器发来的POST请求，用户对应动作触发
func ServicePost(c *gin.Context) {
	tools.GetLogger().Infof("ServicePost url : %v", c.Request.URL.String())
	var input map[string]string
	err := c.ShouldBindXML((*model.StringMap)(&input))
	if err != nil {
		tools.GetLogger().Errorf("bind xml error : %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  constant.StatusCodeMessageMap[constant.StatusCodeInputError],
		})
		return
	}
	tools.GetLogger().Infof("input %v", input)
	c.String(http.StatusOK, "Success")
}
