package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

func SetupDB() {
	db := model.ConnectionDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
