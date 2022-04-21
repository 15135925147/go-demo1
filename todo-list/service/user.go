package service

import (
	"github.com/15135925147/go-demo1/todo-list/model"
	"github.com/15135925147/go-demo1/todo-list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		// 用户已存在
		return serializer.Response{
			Status: 400,
			Msg:    "用户名已存在",
		}
	}
	user.UserName = service.UserName
	// 密码加密 再存入model.user
	if err := user.SetPassword(service.Password); err != nil {
		// 加密失败
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 数据库添加这个用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库添加用户失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}

}
