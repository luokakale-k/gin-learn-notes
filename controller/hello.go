package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + name + "!",
	})
}
