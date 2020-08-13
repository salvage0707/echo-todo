package server

import (
	"github.com/labstack/echo/v4"
)

// Run サーバ起動
func Run(port string) {
	// Echo instance
	e := echo.New()

	// Middleware
	setupMiddleware(e)

	// Routes
	setupRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
