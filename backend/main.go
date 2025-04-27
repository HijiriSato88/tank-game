package main

import (
	"log"

	"backend/db"
	"backend/handler"
	"backend/pkg/jwtutil"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.SetupDB()

	e := echo.New()

	// 新規登録、ログイン
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	// ログイン以降
	auth := e.Group("/auth")
	auth.Use(jwtutil.JWTMiddleware())
	auth.GET("/me", handler.Me)
	auth.POST("/score", handler.UpdateScore)

	// 敵データ取得
	e.GET("/enemies", handler.GetEnemies)
	e.GET("/enemies/name", handler.GetEnemyByNameHandler)


	e.Logger.Fatal(e.Start(":8080"))
}
