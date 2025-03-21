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

func GetUserInfo(c *gin.Context) {
	var req request.GetUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}

	user, err := service.GetUserByID(req.ID)
	if err != nil {
		utils.Fail(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}

func UpdateUser(c *gin.Context) {
	var req request.UpdateUserRequest
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

	err := service.UpdateUser(req)
	if err != nil {
		utils.Fail(c, "用户信息更新失败:"+err.Error())
	}

	utils.Success(c, nil)
}

func DeleteUser(c *gin.Context) {
	var req request.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}

	err := service.DeleteUser(req.ID)
	if err != nil {
		utils.Fail(c, "用户删除失败"+err.Error())
		return
	}
	utils.Success(c, nil)
}

func UserList(c *gin.Context) {
	var req request.UserListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, "参数错误")
		return
	}

	users, total, err := service.GetUserList(req)
	if err != nil {
		utils.Fail(c, "获取用户列表失败:"+err.Error())
		return
	}

	utils.Success(c, utils.PageData{
		List:     users,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}
