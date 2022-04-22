package router

import (
	"github.com/15135925147/go-demo1/todo-list/api"
	"github.com/15135925147/go-demo1/todo-list/middleware"
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
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //每次请求都需要验证
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks", api.ListTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return router
}
