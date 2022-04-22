package service

import (
	"github.com/15135925147/go-demo1/todo-list/model"
	"github.com/15135925147/go-demo1/todo-list/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `form:"title" json:"title"`     //主题
	Content string `form:"content" json:"content"` //内容
	Status  int    `form:"status" json:"status"`   //0是未完成   1是已完成
}

type ShowTaskService struct {
}

// 参数1：uid  参数2：url传来的task id
func (showTask *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	err := model.DB.First(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "获取备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "获取成功",
		Data:   serializer.BuilderTask(task),
	}
}

func (createTask *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     createTask.Title,
		Content:   createTask.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录数据失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建备忘录成功",
	}
}
