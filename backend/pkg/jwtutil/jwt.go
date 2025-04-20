package jwtutil

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))
}

func JWTMiddleware() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")

	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(secret),
		TokenLookup: "header:Authorization",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(CustomClaims)
		},
		ContextKey: "user",
	})
}

func ExtractUser(c echo.Context) (*CustomClaims, error) {
	fmt.Println("ExtractUser called")

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		fmt.Println("user is not *jwt.Token")
		return nil, fmt.Errorf("invalid token format")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		fmt.Printf("claims type mismatch: %#v\n", token.Claims)
		return nil, fmt.Errorf("invalid claims type")
	}

	fmt.Printf("Extracted user_id: %d\n", claims.UserID)
	return claims, nil
}
