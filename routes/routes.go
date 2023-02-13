package routes

import (
	//"go-echo-api/controllers"
	"github.com/labstack/echo/v4"
	"go-echo-api/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/pegawai", controllers.FetchListPegawaiController())

	return e
}
