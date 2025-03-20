package main

import (
	"gin-learn-notes/router"
)

func main() {
	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	r.Run(":8080")
}
