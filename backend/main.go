package main

import (
	"net/http"

	"backend/db"
	"backend/handler"
	"backend/middleare"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Initialize()
	db.SetupDB()
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "backend is go + mysql + redis + docker")
	})

	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	auth := e.Group("/auth")
	auth.Use(middleware.JWT())  // 自作のJWTミドルウェア関数
	auth.GET("/me", handler.Me)

	e.Logger.Fatal(e.Start(":8080"))
}