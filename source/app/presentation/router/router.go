package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func IncludeRouter(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, world!")
	})
}
