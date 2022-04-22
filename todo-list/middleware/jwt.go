package middleware

import (
	"github.com/15135925147/go-demo1/todo-list/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		//fmt.Println("JWT get Authorization:", token)
		if token == "" {
			code = 404 //没有token
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // token 无权限
			} else if time.Now().Unix() > claims.ExpiresAt { //当前时间大于token失效时间，即token以失效
				code = 401 //token已经失效
			}

		}
		if code != 200 {
			c.JSON(http.StatusOK, gin.H{
				"Status": code,
				"Msg":    "Token 验证失败",
			})
			c.Abort() //之后的代码终止运行，直接返回
			return
		}
		c.Next() //挂起该函数  运行之后的函数
	}
}
