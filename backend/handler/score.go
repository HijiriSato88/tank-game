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

func UpdateScore(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}

	var req ScoreRequest
	if err := c.Bind(&req); err != nil || req.Score < 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid score"})
	}

	user, err := repository.GetUserByID(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "user not found"})
	}

	if req.Score > user.BestScore {
		err = repository.UpdateBestScore(user.ID, req.Score)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update score"})
		}
		user.BestScore = req.Score
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":   "score updated",
		"id":        user.ID,
		"username":  user.Username,
		"bestScore": user.BestScore,
		"createdAt": user.CreatedAt,
	})
}
