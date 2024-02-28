package main

import (
	"fmt"
	"os"
	"yoyaku/app/infrastructure/database/postgres"
	"yoyaku/app/presentation/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	// DB初期化
	db := postgres.OpenDB()
	defer postgres.CloseDB(db)

	// ルーターの設定
	router.IncludeRouter(e, db)

	// サーバーの起動
	port := os.Getenv("MY_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
