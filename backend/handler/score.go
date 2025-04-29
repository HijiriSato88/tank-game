package handler

import (
	"backend/pkg/jwtutil"
	"backend/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ScoreRequest struct {
	Score int `json:"score"`
}

func InsertScore(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}

	var req ScoreRequest
	if err := c.Bind(&req); err != nil || req.Score < 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid score"})
	}

	err = repository.InsertScore(claims.UserID, req.Score)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to insert score"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "score inserted successfully"})
}
