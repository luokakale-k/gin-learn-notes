package middleware

import (
	"gin-learn-notes/core/response"
	"gin-learn-notes/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, response.InvalidToken, "Token无效或者不存在")
			c.Abort()
			return
		}

		// 处理 "Bearer xxxxx" 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Fail(c, response.TokenFormatError, "Token格式错误")
			c.Abort()
			return
		}

		// 解析 token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Fail(c, response.TokenExpired, "Token无效或已过期")
			c.Abort()
			return
		}

		// 注入 userID 到上下文
		c.Set("userID", claims.UserID)

		// 放行
		c.Next()
	}
}
