package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PW"), os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		fmt.Println(err)
	}
}
