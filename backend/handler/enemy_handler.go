package handler

import (
	"backend/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEnemies(c echo.Context) error {
	enemies, err := repository.GetAllEnemies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get enemies",
		})
	}

	return c.JSON(http.StatusOK, enemies)
}

func GetEnemyByNameHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "name parameter is required",
		})
	}

	enemy, err := repository.GetEnemyByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "enemy not found",
		})
	}

	return c.JSON(http.StatusOK, enemy)
}
