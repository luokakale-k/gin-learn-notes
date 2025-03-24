package controller

import (
	"errors"
	"gin-learn-notes/core/response"
	"gin-learn-notes/logger"
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
			response.Fail(c, response.ParamError, msg)
		} else {
			// 其他绑定错误，如 JSON 格式错误
			response.Fail(c, response.ParamError, "参数错误")
		}
		return
	}

	user, err := service.RegisterUser(req)
	if err != nil {
		response.Fail(c, response.DBError, "用户保存失败"+err.Error())
	}

	response.Success(c, gin.H{
		"user_id": user.ID,
	})
}

func GetUserInfo(c *gin.Context) {
	var req request.GetUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.ParamError, "参数错误")
		return
	}

	user, err := service.GetUserByID(req.ID)
	if err != nil {
		response.Fail(c, response.NotFound, "用户不存在")
		return
	}

	response.Success(c, user)
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
			response.Fail(c, response.ParamError, msg)
		} else {
			// 其他绑定错误，如 JSON 格式错误
			response.Fail(c, response.ParamError, "参数错误")
		}
		return
	}

	err := service.UpdateUser(req)
	if err != nil {
		response.Fail(c, response.DBError, "用户信息更新失败")
		return
	}

	response.Success(c, nil)
}

func DeleteUser(c *gin.Context) {
	var req request.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.ParamError, "参数错误")
		return
	}

	err := service.DeleteUser(req.ID)
	if err != nil {
		response.Fail(c, response.DBError, "用户删除失败")
		return
	}
	response.Success(c, nil)
}

func UserList(c *gin.Context) {
	var req request.UserListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.ParamError, "参数错误")
		return
	}

	users, total, err := service.GetUserList(req)
	if err != nil {
		response.Fail(c, response.DBError, "获取用户列表失败")
		return
	}

	logger.Log.Info("用户列表：", users)

	response.Success(c, response.PageData{
		List:     users,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
}

func GetUserProfile(c *gin.Context) {
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.Fail(c, response.Unauthorized, "未登录或 token 无效")
		return
	}

	user, err := service.GetUserProfileWithCache(userID)
	if err != nil {
		response.Fail(c, response.NotFound, "用户不存在")
		return
	}

	response.Success(c, user)
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, response.ParamError, "参数错误")
		return
	}

	res, err := service.DoLogin(req.Username, req.Password)
	if err != nil {
		response.Fail(c, response.Unauthorized, "用户名或密码错误")
		return
	}

	response.Success(c, gin.H{
		"token":    res.Token,
		"user_id":  res.UserID,
		"username": res.Username,
	})
}
