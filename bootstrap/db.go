package bootstrap

import (
	"github.com/zhangguangying/goblog/pkg/model"
	"time"
)

func SetDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
