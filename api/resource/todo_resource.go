package resource

import (
	"echo_sample/model"
	"time"
)

// Todo Resource
type Todo struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Finished  *bool     `json:"finished,default:18"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoCreateRequest TodoController#Createのリクエスト
type TodoCreateRequest struct {
	Title string `json:"title" validate:"required,min=1,max=50"`
}

// TodoEditRequest TodoController#Editのリクエスト
type TodoEditRequest struct {
	Title string `json:"title" validate:"required,min=1,max=50"`
}

// TodoChangeFinishedRequest TodoController#ChangeFinishedのリクエスト
type TodoChangeFinishedRequest struct {
	Finished *bool `json:"finished" validate:"required"`
}

// MapToTodoResource modelからresourceに変換する
func MapToTodoResource(m model.Todo) (r Todo) {
	r.ID = m.ID
	r.Title = m.Title
	r.Finished = &m.Finished
	r.CreatedAt = m.CreatedAt

	return
}
