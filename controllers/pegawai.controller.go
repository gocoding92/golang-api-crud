package controllers

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/models"
	"net/http"
)

func FetchListPegawaiController(c echo.Context) error {
	result := models.FetchListPegawaiModel()

	return c.JSON(http.StatusOK, result)
}
