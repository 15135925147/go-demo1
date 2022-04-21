package main

import (
	"fmt"
	Conf "github.com/15135925147/go-demo1/todo-list/conf"
	"github.com/15135925147/go-demo1/todo-list/router"
	"strconv"
)

func main() {
	Conf.GetConf()
	r := router.NewRouter()
	r.Run(fmt.Sprint(":", strconv.Itoa(Conf.Config.Service.HTTPPort)))
	//fmt.Printf("hello %v", "Jayson")
}
