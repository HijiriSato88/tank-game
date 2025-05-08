package router

import (
	"backend/handler"
	"backend/infra"
	"backend/pkg/jwtutil"
	"backend/usecase"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	// enemy
	enemyRepo := infra.NewEnemyRepository()
	enemyUsecase := usecase.NewEnemyUsecase(enemyRepo)
	enemyHandler := handler.NewEnemyHandler(enemyUsecase)

	// ranking
	rankingRepo := infra.NewRankingRepository()
	rankingUsecase := usecase.NewRankingUsecase(rankingRepo)
	rankingHandler := handler.NewRankingHandler(rankingUsecase)

	// user
	userRepo := infra.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo, rankingRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// 公開ルート
	e.POST("/signup", userHandler.Signup)
	e.POST("/login", userHandler.Login)
	e.GET("/enemies", enemyHandler.GetEnemies)
	e.GET("/enemy", enemyHandler.GetEnemyByName)
	e.GET("/ranking", rankingHandler.GetRanking)

	// 認証必要なルート
	auth := e.Group("/auth")
	auth.Use(jwtutil.JWTMiddleware())
	auth.GET("/me", userHandler.Me)
	auth.POST("/score", userHandler.UpdateHighScore)
}
