package main

import (
	"echo_sample/database"
	"echo_sample/server"
	"echo_sample/settings/config"
	"flag"
)

func main() {
	port := flag.String("port", "8000", "サーバポート")
	mode := flag.String("mode", "development", "実行環境")
	flag.Parse()

	config.SetupConfig(*mode)
	database.SetupDatabase()
	defer database.CloseDatabase()

	server.Run(*port)
}
