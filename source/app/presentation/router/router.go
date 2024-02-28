package router

import (
	"net/http"
	"os"
	appmiddleware "yoyaku/app/presentation/middleware"
	"yoyaku/app/service/schema"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func IncludeRouter(e *echo.Echo, db *gorm.DB, logger *zap.Logger) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, schema.DefaultResponseModel{
			Message: "Hello, world!",
		})
	})

	// apiグループ
	api := e.Group("/api")

	// CORSミドルウェアの設定
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken,
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFミドルウェアの設定
	api.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("MY_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieMaxAge:   60,
	}))

	// リクエストミドルウェアの設定
	api.Use(appmiddleware.Request(logger))

	// トランザクションミドルウェアの設定
	api.Use(appmiddleware.Transaction(db))

	// v1グループ
	v1 := api.Group("/v1")

	// JWTミドルウェアの設定
	v1.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_KEY")),
		TokenLookup: "cookie:token",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, schema.DefaultResponseModel{
				Message: err.Error(),
			})
		},
	}))
}
