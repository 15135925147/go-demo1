package Conf

import (
	"fmt"
	"github.com/15135925147/go-demo1/todo-list/model"
	"github.com/spf13/viper"
)

var (
	Conf *viper.Viper

	Config *model.Config
)

func GetConf() {
	Conf = viper.New()
	Conf.AddConfigPath(".")
	Conf.SetConfigName("config")
	Conf.SetConfigType("yaml")
	if err := Conf.ReadInConfig(); err != nil {
		panic(err)
	}
	loadConf()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=true", Config.Mysql.DBUser, Config.Mysql.Dbpw, Config.Mysql.DBAddr, Config.Mysql.DBPort, Config.Mysql.DBName)
	model.InitDB(dsn)
}

func loadConf() {
	Config = &model.Config{}
	// load service
	Config.Service.AppModel = Conf.GetString("service.AppModel")
	Config.Service.HTTPPort = Conf.GetInt("service.HttpPort")

	// load redis
	Config.Redis.RedisAddr = Conf.GetString("redis.RedisAddr")
	Config.Redis.RedisPort = Conf.GetInt("redis.RedisPort")
	Config.Redis.RedisPW = Conf.GetString("redis.RedisPW")
	Config.Redis.RedisDBName = Conf.GetInt("redis.RedisDBName")

	// load mysql
	Config.Mysql.DBDrive = Conf.GetString("Mysql.DBDrive")
	Config.Mysql.DBAddr = Conf.GetString("Mysql.DBAddr")
	Config.Mysql.DBPort = Conf.GetInt("Mysql.DBPort")
	Config.Mysql.DBUser = Conf.GetString("Mysql.DBUser")
	Config.Mysql.Dbpw = Conf.GetString("Mysql.DBPW")
	Config.Mysql.DBName = Conf.GetString("Mysql.DBName")

}
