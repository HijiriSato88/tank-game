package handler

import (
	"backend/pkg/jwtutil"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ScoreHandler struct {
	scoreUsecase usecase.ScoreUsecase
}

func NewScoreHandler(u usecase.ScoreUsecase) *ScoreHandler {
	return &ScoreHandler{scoreUsecase: u}
}

type ScoreRequest struct {
	Score int `json:"score"`
}

func (h *ScoreHandler) InsertScore(c echo.Context) error {
	claims, err := jwtutil.ExtractUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}

	var req ScoreRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	err = h.scoreUsecase.InsertScore(claims.UserID, req.Score)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to insert score"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "score inserted successfully"})
}
