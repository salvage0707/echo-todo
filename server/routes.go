package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupRouter(e *echo.Echo) {
	e.GET("/", hello)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!desu")
}
