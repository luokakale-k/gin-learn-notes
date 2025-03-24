package router

import (
	"gin-learn-notes/controller"
	"gin-learn-notes/middleware"
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

	r.POST("/login", controller.Login)

	// 需要登录的接口分组
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.POST("/profile", controller.GetUserProfile)
	}
	return r
}
