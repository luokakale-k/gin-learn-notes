package utils

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) uint {
	if userIDRaw, exists := c.Get("userID"); exists {
		if userID, ok := userIDRaw.(uint); ok {
			return userID
		}
	}
	return 0
}
