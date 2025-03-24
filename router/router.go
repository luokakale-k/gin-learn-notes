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

	r.POST("/info", controller.GetUserInfo)

	r.POST("/save", controller.UpdateUser)

	r.POST("/delete", controller.DeleteUser)

	r.POST("/list", controller.UserList)

	r.POST("/profile", controller.GetUserProfile)

	return r
}
