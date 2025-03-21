package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": -1,
		"msg":  msg,
	})
}
