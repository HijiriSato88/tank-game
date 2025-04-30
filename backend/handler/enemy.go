package handler

import (
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EnemyHandler struct {
	enemyUsecase usecase.EnemyUsecase
}

func NewEnemyHandler(u usecase.EnemyUsecase) *EnemyHandler {
	return &EnemyHandler{enemyUsecase: u}
}

func (h *EnemyHandler) GetEnemyByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "name parameter is required",
		})
	}

	enemy, err := h.enemyUsecase.GetByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "enemy not found",
		})
	}

	return c.JSON(http.StatusOK, enemy)
}

func (h *EnemyHandler) GetEnemies(c echo.Context) error {
	enemies, err := h.enemyUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get enemies",
		})
	}
	return c.JSON(http.StatusOK, enemies)
}