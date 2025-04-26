package handler

import (
	"backend/pkg/jwtutil"
	"backend/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Me(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}

	u, err := repository.GetUserByID(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"username":  u.Username,
	})
}
