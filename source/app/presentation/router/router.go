package router

import (
	"net/http"
	"yoyaku/app/service/schema"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func IncludeRouter(e *echo.Echo, db *gorm.DB, logger *zap.Logger) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, schema.DefaultResponseModel{
			Message: "Hello, world!",
		})
	})
}
