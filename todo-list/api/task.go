package api

import (
	"github.com/15135925147/go-demo1/todo-list/pkg/utils"
	"github.com/15135925147/go-demo1/todo-list/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 创建备忘录
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	if err := c.ShouldBind(&createTask); err != nil {
		// 绑定失败
		logging.Error(err)
		c.JSON(400, gin.H{
			"err": err,
		})
	} else {
		//验证token
		claims, err := utils.ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(403, gin.H{
				"err": err,
			})
		}

		// 绑定成功  根据user id执行创建
		res := createTask.Create(claims.Id)
		c.JSON(200, res)
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err != nil {
		// 绑定失败
		logging.Error(err)
		c.JSON(400, gin.H{
			"err": err,
		})
	} else {

		// 绑定成功  根据url id 获取到这条备忘录
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	}
}
