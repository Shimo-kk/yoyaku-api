package controller

import (
	"net/http"
	"yoyaku/app/service/schema"

	"github.com/labstack/echo/v4"
)

type authController struct {
}

// 認証コントローラーの作成
func NewAuthController() *authController {
	return &authController{}
}

// CSRFトークンの取得
func (ac *authController) GetCsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)

	resonseBody := schema.CSRFModel{
		Csrf: token,
	}
	return c.JSON(http.StatusOK, resonseBody)
}
