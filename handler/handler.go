package handler

import (
	"net/http"
	"strconv"

	"menu/controller"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return err
	}
	ms, err := controller.GetProduct(uint(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}

func GetAllProduct(c echo.Context) error {
	ms, err := controller.GetAllProduct()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ms)
}
