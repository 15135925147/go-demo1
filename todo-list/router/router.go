package router

import (
	"github.com/15135925147/go-demo1/todo-list/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	router.Use(sessions.Sessions("mysession", store))
	v1 := router.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
	}
	return router
}
