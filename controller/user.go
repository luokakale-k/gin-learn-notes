package controller

import (
	"errors"
	"gin-learn-notes/request"
	"gin-learn-notes/service"
	"gin-learn-notes/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 使用 validator 类型断言
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fieldMap := map[string]string{
				"Name": "用户名",
				"Age":  "年龄",
			}
			msg := utils.TranslateValidationError(ve, fieldMap)
			utils.Fail(c, msg)
		} else {
			// 其他绑定错误，如 JSON 格式错误
			utils.Fail(c, "参数格式不正确")
		}
		return
	}

	user, err := service.RegisterUser(req)
	if err != nil {
		utils.Fail(c, "保存用户失败:"+err.Error())
	}

	utils.Success(c, gin.H{
		"user_id": user.ID,
	})
}
