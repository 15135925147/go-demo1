package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	DB *gorm.DB
)

func InitDB(connString string) {
	DB, err := gorm.Open("mysql", connString)
	if err != nil {
		panic("mysql init failed;")
	}
	fmt.Println("mysql init success；")
	//默认情况下允许输出db的log
	DB.LogMode(true)
	//当gin框架运行发行版的时候不输出db的log
	if gin.Mode() == "release" {
		DB.LogMode(false)
	}

	//对db连接池做配置
	DB.SingularTable(true) // gorm 创建表的时候默认会在最后加s，当为true不会自动加s
	//在池中允许更多的空闲连接将提高性能，因为这样可以减少从头开始建立新连接的可能性，从而有助于节省资源。
	//但是保持空闲连接的存活是要付出代价的——它会占用内存
	DB.DB().SetMaxIdleConns(5) //设置最多保留5个空闲连接
	//如果5个连接都已经打开被使用，并且应用程序需要另一个连接的话，那么应用程序将被迫等待，直到5个打开的连接其中的一个被释放并变为空闲。
	//标准 （实例数）*（最大并发连接数）< (mysql最大连接数)
	DB.DB().SetMaxOpenConns(20)                  //设置最大并发连接数
	DB.DB().SetConnMaxLifetime(time.Second * 30) //连接可重用的最大时间长度
}
