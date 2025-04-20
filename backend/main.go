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

	e.Logger.Fatal(e.Start(":8080"))
}
