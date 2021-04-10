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
	user := api.Group("/user")
	{
		user.POST("/login", handler.UserLogin)
		user.GET("/login/verification_code", handler.VerificationCode)
		//微信跳转获取用户openID接口
		user.GET("/open_id/:code", handler.GetUserOpenIdByCode)
		user.GET("/info", middleware.Auth, handler.GetUserInfo)
		user.POST("/update", middleware.Auth, handler.UpdateUserInfo)
	}
	examination := api.Group("/examination")
	{
		examination.POST("/buy", middleware.Auth, handler.BuyExamination)
		examination.POST("/refund", middleware.Auth, handler.RefundExamination)
		examination.GET("/info", middleware.Auth, handler.GetExaminationInfo)
	}
	calendar := api.Group("/calendar")
	{
		calendar.GET("/list", middleware.Auth, handler.ListExpenseCalendar)
	}
	checking := api.Group("/check")
	{
		checking.POST("/booking", middleware.Auth, handler.BookingCheck)
	}
}
