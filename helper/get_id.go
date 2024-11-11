package helper

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIDParam(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetPlantIDParam(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("plant_id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}
