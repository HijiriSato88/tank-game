package main

import (
	"net/http"

	"backend/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Initialize()
	db.SetupDB()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "backend is go + mysql + redis + docker")
	})

	// DB接続確認エンドポイント
	e.GET("/db-check", func(c echo.Context) error {
		err := db.DB.Ping()
		if err != nil {
			return c.String(http.StatusInternalServerError, "DB接続に失敗しました")
		}
		return c.String(http.StatusOK, "DB接続に成功しました！")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
