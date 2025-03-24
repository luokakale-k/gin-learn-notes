package service

import (
	"encoding/json"
	"fmt"
	"gin-learn-notes/config"
	"gin-learn-notes/model"
	"gin-learn-notes/request"
	"gin-learn-notes/response"
	"gin-learn-notes/utils"
	"strings"
	"time"
)

func RegisterUser(req request.RegisterRequest) (*model.User, error) {
	user := &model.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// 根据 ID 获取用户信息
func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// 更新用户信息
func UpdateUser(req request.UpdateUserRequest) error {
	var user model.User
	if err := config.DB.First(&user, req.ID).Error; err != nil {
		return err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Age != nil {
		user.Age = *req.Age
	}

	return config.DB.Save(&user).Error
}

// 删除用户
func DeleteUser(id uint) error {
	var user model.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		return err
	}

	return config.DB.Delete(&user).Error
}

// 获取用户列表
func GetUserList(req request.UserListRequest) ([]model.User, int64, error) {

	db := config.DB.Model(&model.User{})

	req.Keyword = strings.TrimSpace(req.Keyword)

	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}

	if req.MinAge > 0 {
		db = db.Where("age >= ?", req.MinAge)
	}

	if req.MaxAge > 0 {
		db = db.Where("age <= ?", req.MaxAge)
	}

	return utils.Paginate[model.User](db, req.Page, req.PageSize)
}

// 使用缓存获取用户信息
func GetUserProfileWithCache(userID uint) (*model.User, error) {
	cacheKey := fmt.Sprintf("user:info:%d", userID)

	// 先查询缓存
	cacheStr, err := utils.RedisGet(cacheKey)
	if err == nil && cacheStr != "" {
		var user model.User
		json.Unmarshal([]byte(cacheStr), &user)
		return &user, nil
	}

	// 缓存未命中，查询数据库
	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// 将查询结果写入缓存
	userStr, _ := json.Marshal(user)
	utils.RedisSet(cacheKey, string(userStr), 10*time.Minute)

	return user, nil
}

func DoLogin(username, password string) (*response.LoginResult, error) {
	var user model.User
	if err := config.DB.Where("name= ? AND password= ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}

	// 调用封装的 token 生成方法
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	// 返回 token 和用户信息
	return &response.LoginResult{
		Token:    token,
		UserID:   user.ID,
		Username: user.Name,
	}, nil
}
