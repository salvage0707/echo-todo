package server

import (
	"echo_sample/settings/config"

	"github.com/labstack/echo/v4"
)

// Run サーバ起動
func Run(port string) {
	// Echo instance
	e := echo.New()
	e.Debug = config.Config().GetBool("app.debug")

	// 設定
	SetupValidater(e)
	setupMiddleware(e)
	setupRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
