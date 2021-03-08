package main

import (
	"github.com/MisakiFx/martin/biz/handler"
	"github.com/MisakiFx/martin/biz/middleware"
	"github.com/gin-gonic/gin"
)

func customizeRegister(r *gin.Engine) {
	r.Use(gin.Recovery())

	r.GET("/", middleware.CheckToken, handler.ServiceGet)
	r.POST("/", middleware.CheckToken, handler.ServicePost)

	api := r.Group("/guardian/api")
	{
		api.POST("/login", handler.UserLogin)
	}
}
