package main

import (
	"gin-learn-notes/config"
	"gin-learn-notes/logger"
	"gin-learn-notes/model"
	"gin-learn-notes/router"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 初始化日志
	logger.InitLogger()
	defer logger.Log.Sync()

	// 自动建表
	config.DB.AutoMigrate(&model.User{})

	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	r.Run(":8080")
}
