package routes

import (
	"github.com/labstack/echo/v4"
	"go-echo-api/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/pegawai", controllers.FetchAllPegawaiController)

	return e
}
