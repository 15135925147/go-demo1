package serializer

import "github.com/15135925147/go-demo1/todo-list/model"

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title"`   //标题
	Status    int    `json:"status"`  //0是未完成，1是已完成
	Content   string `json:"content"` //内容
	View      uint64 `json:"view"`    //浏览量
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"` //备忘录开始时间
	EndTime   int64  `json:"end_time"`   //备忘录完成时间

}

func BuilderTask(task model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Status:    task.Status,
		Content:   task.Content,
		CreateAt:  task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}

}

func BuilderTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuilderTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
