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

func (h *EnemyHandler) GetEnemies(c echo.Context) error {
	enemies, err := h.enemyUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get enemies",
		})
	}
	return c.JSON(http.StatusOK, enemies)
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
			"error": "enemy not found in mysql",
		})
	}

	return c.JSON(http.StatusOK, enemy)
}

/*
func GetEnemyByNameFromRedisHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "name parameter is required",
		})
	}

	// Redisから取得（まだ文字列のまま）
	jsonData, err := repository.GetEnemyByNameFromRedis(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "enemy not found in redis",
		})
	}

	// Jsonへ変換
	var enemy model.Enemy
	if err := json.Unmarshal([]byte(jsonData), &enemy); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to parse enemy data",
		})
	}

	return c.JSON(http.StatusOK, enemy)
}
*/