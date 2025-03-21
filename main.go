package main

import (
	"gin-learn-notes/config"
	"gin-learn-notes/model"
	"gin-learn-notes/router"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 自动建表
	config.DB.AutoMigrate(&model.User{})

	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	r.Run(":8080")
}
