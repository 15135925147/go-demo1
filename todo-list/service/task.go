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

type ShowTaskService struct{}

type ListTaskService struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}

type UpdateTaskService struct {
	Title   string `form:"title" json:"title"`     //主题
	Content string `form:"content" json:"content"` //内容
	Status  int    `form:"status" json:"status"`   //0是未完成   1是已完成
}

type DeleteTaskService struct{}

func (listTask *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if listTask.PageSize == 0 {
		listTask.PageSize = 10
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(listTask.PageSize).Offset((listTask.PageNum - 1) * listTask.PageNum).Find(&tasks)
	return serializer.BuilderListResponse(serializer.BuilderTasks(tasks), count)
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

func (updateTask *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	code := 200
	model.DB.First(&task, tid)

	task.Title = updateTask.Title
	task.Content = updateTask.Content
	task.Status = updateTask.Status

	err := model.DB.Save(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "更新备忘录数据失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuilderTask(task),
		Msg:    "更新备忘录成功",
	}
}

func (deleteTask *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "删除备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
