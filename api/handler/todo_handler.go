package handler

import (
	"echo_sample/api/resource"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// TodoHandler Todoハンドラ
type TodoHandler struct {
	db *gorm.DB
}

// Create todoを登録する
func (h TodoHandler) Create(c echo.Context) (err error) {
	todoRequest := new(resource.TodoCreateRequest)
	if err = c.Bind(todoRequest); err != nil {
		return
	}

	// todo := model.Todo{}
	// todo.Title = todoRequest.Title

	// createdTodoModel := ctrl.todoService.Create(todo)

	// createdTodo := copyModel(createdTodoModel)

	// c.JSON(http.StatusCreated, createdTodo)
}
