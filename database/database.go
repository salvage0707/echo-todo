package database

import (
	"echo_sample/model"
	"echo_sample/settings/config"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

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
	connectionURL, err := pq.ParseURL(dbURL)
	if err != nil {
		log.Error(err.Error())
		panic(fmt.Sprintf("データベースURLの解析に失敗しました。 url: %s", dbURL))
	}

	if config.IsDevelopmentMode() {
		connectionURL += " sslmode=disable"
	} else {
		connectionURL += " sslmode=require"
	}

	db, err := gorm.Open("postgres", connectionURL)
	if err != nil {
		log.Error(err.Error())
		panic(fmt.Sprintf("データベース接続に失敗しました。 url: %s", dbURL))
	}

	return db
}

func autoMigration() {
	// スキーマのマイグレーション
	dbConnection.AutoMigrate(&model.Todo{})
}
