package handler

import (
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RankingHandler struct {
	rankingUsecase usecase.RankingUsecase
}

func NewRankingHandler(u usecase.RankingUsecase) *RankingHandler {
	return &RankingHandler{rankingUsecase: u}
}

func (h *RankingHandler) GetRanking(c echo.Context) error {
	limitParam := c.QueryParam("limit")
	limit := 10
	if limitParam != "" {
		if parsed, err := strconv.Atoi(limitParam); err == nil {
			limit = parsed
		}
	}
	if limit <= 0 || limit > 10000 {
		limit = 1000
	}

	rankings, err := h.rankingUsecase.GetRanking(limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get ranking"})
	}

	return c.JSON(http.StatusOK, rankings)
}
