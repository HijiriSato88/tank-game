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

	// DB設定
	db.SetupDB()
	db.InitRedis()

	// echo設定
	e := echo.New()

	// user
	userRepo := infra.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// score
	scoreRepo := infra.NewScoreRepository()
	scoreUsecase := usecase.NewScoreUsecase(scoreRepo)
	scoreHandler := handler.NewScoreHandler(scoreUsecase)

	// 新規登録、ログイン
	e.POST("/signup", userHandler.Signup)
	e.POST("/login", userHandler.Login)

	// ログイン以降
	auth := e.Group("/auth")
	auth.Use(jwtutil.JWTMiddleware())
	auth.GET("/me", userHandler.Me)
	auth.POST("/score", scoreHandler.InsertScore)

	// 敵データ取得
	e.GET("/enemies", handler.GetEnemies)
	e.GET("/enemies/name", handler.GetEnemyByNameFromRedisHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
