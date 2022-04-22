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
		//这里是为了获取到用户信息，  主要不是为了验证token因为进入这里之前中间件已经验证了token
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

		// 绑定成功  根据user id执行创建
		res := createTask.Create(claims.Id)
		c.JSON(200, res)
	}
}

//查看某一条备忘录
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

//获取到该用户所有的备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	if err := c.ShouldBind(&listTask); err != nil {
		// 绑定失败
		logging.Error(err)
		c.JSON(400, gin.H{
			"err": err,
		})
	} else {
		//这里是为了获取到用户信息，  主要不是为了验证token因为进入这里之前中间件已经验证了token
		claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

		// 绑定成功  根据user id获取到该用户所有的备忘录
		res := listTask.List(claims.Id)
		c.JSON(200, res)
	}
}

//更新备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err != nil {
		// 绑定失败
		logging.Error(err)
		c.JSON(400, gin.H{
			"err": err,
		})
	} else {
		// 绑定成功  根据user id执行创建
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

//删除备忘录
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	// 绑定成功  根据url穿进来id执行删除
	res := deleteTask.Delete(c.Param("id"))
	c.JSON(200, res)

}
