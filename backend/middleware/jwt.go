package middleware

import (
	"os"
	"github.com/labstack/echo/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func JWT() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "header:Authorization",
		ContextKey:  "user",
	})
}