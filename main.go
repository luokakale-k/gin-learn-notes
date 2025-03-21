package main

import (
	"fmt"
	"gin-learn-notes/config"
	"gin-learn-notes/logger"
	"gin-learn-notes/router"
)

func main() {

	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	config.InitDB()

	// 初始化日志
	logger.InitLogger()
	defer logger.Log.Sync()

	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	addr := fmt.Sprintf(":%d", config.Conf.App.Port)
	r.Run(addr)
}
