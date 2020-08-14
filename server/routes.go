package server

import (
	"echo_sample/api/handler"
	"echo_sample/database"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func setupRouter(e *echo.Echo) {
	db := database.GetConnection()

	{
		g := e.Group("/todos")
		h := &handler.TodoHandler{DB: db}
		g.GET("", h.Index)
		g.GET("/:id", h.Show)
		g.POST("", h.Create)
		g.PATCH("/:id", h.Edit)
		g.PUT("/:id", h.ChangeFinished)
		g.DELETE("/:id", h.Delete)
	}

	exportRoutes(e)
}

func exportRoutes(e *echo.Echo) {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Debug(err.Error())
	}
	ioutil.WriteFile("routes.json", data, 0644)
}
