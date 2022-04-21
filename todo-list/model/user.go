package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string //保存的是加密后的密码
}

// 传入的密码进行加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// 传入的密码进行解密
func (user *User) CheckPassword(password string) bool {
	// 将传入的密码 跟数据库的密码进行比较
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))

	return err == nil
}
