package main

import (
	"net/http"

	"github.com/MisakiFx/martin/pkg/connection"

	"github.com/gin-gonic/gin"
)

func Init() {
	//tools.Init()
	connection.Init()
}

func main() {
	Init()

	r := gin.Default()

	customizeRegister(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pong": "pong",
		})
	})

	r.Run(":8080")
}
