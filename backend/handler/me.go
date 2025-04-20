package handler

import (
	"net/http"
	"backend/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Me(c echo.Context) error {
	// JWTトークンのclaimsを取り出す
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int(claims["user_id"].(float64))

	// DBからユーザー取得
	u, err := repository.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "user not found"})
	}

	// 必要な情報だけ返す（パスワードは除外）
	return c.JSON(http.StatusOK, echo.Map{
		"id":        u.ID,
		"username":  u.Username,
		"bestScore": u.BestScore,
		"createdAt": u.CreatedAt,
	})
}
