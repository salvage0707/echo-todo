package database

import (
	"echo_sample/model"
	"echo_sample/settings/config"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	// SQLite
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConnection *gorm.DB

// SetupDatabase DBセットアップ
func SetupDatabase() {
	dbConnection = initConnection()

	if config.Config().GetBool("app.debug") {
		dbConnection.Debug()
		dbConnection.LogMode(true)
	}

	// SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定します
	dbConnection.DB().SetMaxIdleConns(10)
	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	dbConnection.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	dbConnection.DB().SetConnMaxLifetime(time.Hour)

	autoMigration()
}

// CloseDatabase DBコネクションをクローズする
func CloseDatabase() {
	dbConnection.Close()
}

// GetConnection DBコネクションを取得する
func GetConnection() *gorm.DB {
	return dbConnection
}

func initConnection() *gorm.DB {
	dbURL := config.Config().GetString("database.url")

	db, err := gorm.Open("postgres", dbURL+" sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("データベースへの接続に失敗しました。 url: %s", dbURL))
	}

	return db
}

func autoMigration() {
	// スキーマのマイグレーション
	dbConnection.AutoMigrate(&model.Todo{})
}
