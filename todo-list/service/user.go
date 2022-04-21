package service

import (
	"fmt"
	"github.com/15135925147/go-demo1/todo-list/model"
	"github.com/15135925147/go-demo1/todo-list/pkg/utils"
	"github.com/15135925147/go-demo1/todo-list/serializer"
	"github.com/jinzhu/gorm"
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

func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//err是否是用户不存在
		if gorm.ErrRecordNotFound == err {
			return serializer.Response{
				Status: 400,
				Msg:    "该用户不存在",
			}
		}
		//其他错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "CheckPassword:密码错误",
		}
	}
	//登陆成功返回一个token给用户  方便之后功能使用
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	fmt.Println("token:", token)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token 签发失败",
			Error:  err.Error(),
		}
	}
	// 返回token
	return serializer.Response{
		Status: 200,
		Msg:    "登陆成功",
		Data:   serializer.TokenData{User: serializer.Builder(user), Token: token},
	}
}
