package handler

import (
	"echo_sample/api/resource"
	"echo_sample/model"
	"echo_sample/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// TodoHandler Todoハンドラ
type TodoHandler struct {
	DB *gorm.DB
}

// Index todoのリストを返す
func (h TodoHandler) Index(c echo.Context) (err error) {

	tx := h.DB.Model(&model.Todo{})
	// 最終ID
	lastID := c.QueryParam("lastId")
	if lastID != "" {
		tx = tx.Where("id < ?", lastID)
	}

	var todos []model.Todo
	tx.Order("created_at desc").Limit(10).Find(&todos)

	var createdTodos []resource.Todo
	for _, todo := range todos {
		tr := resource.MapToTodoResource(todo)
		createdTodos = append(createdTodos, tr)
	}

	return c.JSON(http.StatusOK, createdTodos)
}

// Show todoを取得する
func (h TodoHandler) Show(c echo.Context) (err error) {
	id := c.Param("id")
	id64, _ := strconv.ParseUint(id, 10, 64)

	todo := model.Todo{}
	todo.ID = id64
	if h.DB.First(&todo).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	findedTodo := resource.MapToTodoResource(todo)
	return c.JSON(http.StatusOK, findedTodo)
}

// Create todoを登録する
func (h TodoHandler) Create(c echo.Context) (err error) {
	todoRequest := new(resource.TodoCreateRequest)
	if err = c.Bind(todoRequest); err != nil {
		return
	}

	if ve := c.Validate(todoRequest); ve != nil {
		ver := utils.ValidateResponse(ve.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, ver)
		return
	}

	todo := model.Todo{}
	todo.Title = todoRequest.Title

	h.DB.Create(&todo)

	createdTodo := resource.MapToTodoResource(todo)

	return c.JSON(http.StatusCreated, createdTodo)
}

// Edit todoを編集する
func (h TodoHandler) Edit(c echo.Context) (err error) {
	id := c.Param("id")
	id64, _ := strconv.ParseUint(id, 10, 64)

	todoRequest := new(resource.TodoEditRequest)
	if err = c.Bind(todoRequest); err != nil {
		return
	}

	if ve := c.Validate(todoRequest); ve != nil {
		ver := utils.ValidateResponse(ve.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, ver)
		return
	}

	todo := model.Todo{}
	todo.ID = id64
	todo.Title = todoRequest.Title

	var count int
	h.DB.Model(&model.Todo{}).Where("id = ?", todo.ID).Count(&count)
	if count == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	h.DB.Model(&todo).Updates(todo)

	editedTodo := resource.MapToTodoResource(todo)

	return c.JSON(http.StatusOK, editedTodo)
}

// ChangeFinished 完了・未完了を変更する
func (h TodoHandler) ChangeFinished(c echo.Context) (err error) {
	id := c.Param("id")
	id64, _ := strconv.ParseUint(id, 10, 64)

	todoRequest := new(resource.TodoChangeFinishedRequest)
	if err = c.Bind(todoRequest); err != nil {
		return
	}

	if ve := c.Validate(todoRequest); ve != nil {
		ver := utils.ValidateResponse(ve.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, ver)
		return
	}

	todo := model.Todo{}
	todo.ID = id64

	var count int
	h.DB.Model(&model.Todo{}).Where("id = ?", todo.ID).Count(&count)
	if count == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	h.DB.Model(&todo).Update("finished", *todoRequest.Finished).Find(&todo)

	editedTodo := resource.MapToTodoResource(todo)

	return c.JSON(http.StatusOK, editedTodo)
}

// Delete Todoを論理削除
func (h TodoHandler) Delete(c echo.Context) (err error) {
	id := c.Param("id")
	id64, _ := strconv.ParseUint(id, 10, 64)

	var count int
	h.DB.Model(&model.Todo{}).Where("id = ?", id64).Count(&count)
	if count == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	h.DB.Where("id = ?", id64).Delete(&model.Todo{})

	var response struct {
		Message string `json:"message"`
	}
	response.Message = fmt.Sprintf("deleted todo. id=%v", id64)

	return c.JSON(http.StatusOK, response)
}
