package api

import (
	"fmt"
	"github.com/15135925147/go-demo1/todo-list/service"
	"github.com/gin-gonic/gin"
)

// 用户注册
func UserRegister(c *gin.Context) {
	// 用户注册的信息的记录
	var userRegister service.UserService
	// 将用户输入的信息绑定的上面用于记录的位置
	if err := c.ShouldBind(&userRegister); err != nil {
		// 绑定失败
		c.JSON(400, nil)
		fmt.Println(err)
	} else {
		// 绑定成功 service执行注册 给用户返回一个标准返回
		res := userRegister.Register()
		c.JSON(200, res)
	}
}
