package controllers

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/models"
	"net/http"
)

func FetchAllPegawaiController(c echo.Context) error {
	result := models.FetchAllPegawaiModel()

	return c.JSON(http.StatusOK, result)
}
