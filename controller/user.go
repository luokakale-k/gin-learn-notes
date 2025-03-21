package controller

import (
	"gin-learn-notes/config"
	"gin-learn-notes/model"
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

	user := model.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
		"user_id": user.ID,
	})
}
