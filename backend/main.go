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

	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	auth := e.Group("/auth")
	auth.Use(jwtutil.JWTMiddleware())
	auth.GET("/me", handler.Me)

	e.Logger.Fatal(e.Start(":8080"))
}
