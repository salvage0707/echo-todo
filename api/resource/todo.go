package resource

import (
	"time"
)

// Todo Resource
type Todo struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Finished  bool      `json:"finished"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoCreateRequest TodoController#Createのリクエスト
type TodoCreateRequest struct {
	Title string `json:"title" binding:"required,min=1,max=50"`
}

// TodoEditRequest TodoController#Editのリクエスト
type TodoEditRequest struct {
	Title string `json:"title" binding:"required,min=1,max=50"`
}
