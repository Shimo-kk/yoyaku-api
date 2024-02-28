package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func IncludeRouter(e *echo.Echo, db *gorm.DB, logger *zap.Logger) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, world!")
	})
}
