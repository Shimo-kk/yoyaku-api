package middleware

import (
	"net/http"
	"yoyaku/app/core"
	"yoyaku/app/service/schema"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Request(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			// アクセスログ
			logger.Info("Access Start",
				zap.String("Host", req.Host),
				zap.String("Protocl", req.Proto),
				zap.String("Path", req.URL.Path),
				zap.String("Method", req.Method),
			)

			// 次の処理を実行
			err := next(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, schema.DefaultResponseModel{
					Message: err.Error(),
				})
			}

			// エラーログ
			if err != nil {
				dstErr := core.AsAppError(err)
				msg := dstErr.Error()
				logger.Error(msg)
			}

			// レスポンスログ
			logger.Info("Access End", zap.Int("Status", res.Status))

			return nil
		}
	}
}
