package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if cookie, err := context.Cookie("cookie"); err == nil && cookie == "cookiezz" {

			context.Next() // 获取成功 进入程序
			return
		} else {
			//
			context.JSON(http.StatusUnauthorized, gin.H{
				"err": err,
			})
			context.Abort() //如果验证失败就不再执行程序，一定要加，不加就等于还要在进入程序
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(context *gin.Context) {
		context.SetCookie("cookie", "cookiezz", 60, "/",
			"localhost", false, true)
		context.String(http.StatusOK, "登陆成功，cookie set success")
	})

	r.GET("/home", AuthMiddleWare(), func(context *gin.Context) {
		cookie, _ := context.Cookie("cookie")
		context.JSON(http.StatusOK, gin.H{
			"path": "home",
			"data": cookie,
		})
	})
	r.Run(":8081")
}
