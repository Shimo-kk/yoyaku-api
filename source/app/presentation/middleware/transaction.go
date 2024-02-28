package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// トランザクションの開始
			tx := db.Begin()
			if tx.Error != nil {
				return errors.New("トランザクションの開始に失敗しました。:" + tx.Error.Error())
			}

			// コンテキストにトランザクションを保存する
			c.Set("tx", tx)

			// 次の処理を実行
			if err := next(c); err != nil {
				tx.Rollback()
				return err
			}

			tx.Commit()
			return nil
		}
	}
}
