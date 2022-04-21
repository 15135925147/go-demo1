package serializer

import "github.com/15135925147/go-demo1/todo-list/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                    //用户ID
	UserName string `json:"user_name" form:"user_name" example:"Jayson"` //用户名
	Status   string `json:"status" form:"status"`                        //用户状态
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func Builder(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
