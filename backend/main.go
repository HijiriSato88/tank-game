package main

import (
	"log"

	"backend/db"
	"backend/handler"
	"backend/usecase"
	"backend/infra"
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
	db.InitRedis()

	e := echo.New()

	r := infra.NewUserRepository()
	u := usecase.NewUserUsecase(r)
	h := handler.NewUserHandler(u)

	// 新規登録、ログイン
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)

	// ログイン以降
	auth := e.Group("/auth")
	auth.Use(jwtutil.JWTMiddleware())
	auth.GET("/me", h.Me)
	auth.POST("/score", handler.InsertScore)

	// 敵データ取得
	e.GET("/enemies", handler.GetEnemies)
	e.GET("/enemies/name", handler.GetEnemyByNameFromRedisHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
