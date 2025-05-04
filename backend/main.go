package main

import (
	"backend/db"
	"backend/router"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	initializeDB()

	e := echo.New()
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initializeDB() {
	db.SetupDB()
	db.InitRedis()
}
