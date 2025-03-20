package router

import (
	"gin-learn-notes/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controller.Ping)

	r.GET("/hello", controller.HelloHandler)

	r.POST("/register", controller.Register)

	return r
}
