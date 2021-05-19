package main

import (
	"net/http"

	"github.com/MisakiFx/martin/martin/pkg/constant"

	"github.com/MisakiFx/martin/martin/biz/handler"
	"github.com/MisakiFx/martin/martin/biz/middleware"
	"github.com/gin-gonic/gin"
)

func customizeRegister(r *gin.Engine) {
	r.Use(gin.Recovery())

	r.Use(handler.CORS)
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
	examination.Use(middleware.Auth)
	{
		examination.POST("/buy", handler.BuyExamination)
		examination.POST("/refund", handler.RefundExamination)
		examination.GET("/info", handler.GetExaminationInfo)
	}
	calendar := api.Group("/calendar")
	calendar.Use(middleware.Auth)
	{
		calendar.GET("/list", handler.ListExpenseCalendar)
	}
	checking := api.Group("/check")
	checking.Use(middleware.Auth)
	{
		checking.POST("/booking", handler.BookingCheck)
		checking.POST("/cancel", handler.CancelBookingCheck)
		checking.GET("/list", handler.ListCheck)
		checking.GET("/result/:id", handler.GetCheckResult)
	}
	admin := api.Group("/admin")
	admin.Use(middleware.Auth, middleware.AdminAuth)
	{
		admin.GET("/", handler.CheckAdmin)
		admin.POST("/check_start", handler.CheckStart)
		admin.POST("/check_finish", handler.CheckFinish)
		admin.POST("/check_result", handler.CheckResult)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": constant.StatusCodeInputError,
			"msg":  "未定义的路由",
		})
		return
	})
}
