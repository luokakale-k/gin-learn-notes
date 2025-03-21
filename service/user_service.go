package service

import (
	"gin-learn-notes/config"
	"gin-learn-notes/model"
	"gin-learn-notes/request"
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
