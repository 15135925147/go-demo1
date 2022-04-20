package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string //保存的是加密后的密码
}
