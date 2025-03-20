package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误：" + err.Error(),
		})
		return
	}

	// 返回欢迎消息
	c.JSON(http.StatusOK, gin.H{
		"message": "欢迎 " + req.Name,
		"age":     req.Age,
	})
}
