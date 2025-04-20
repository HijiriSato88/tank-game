package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func SetupDB() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		dbUser, dbPassword, dbHost, dbName)

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB接続失敗:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB応答なし:", err)
	}
}
