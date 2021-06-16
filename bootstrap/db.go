package bootstrap

import (
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/app/models/category"
	"github.com/zhangguangying/goblog/app/models/user"
	"github.com/zhangguangying/goblog/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migrate(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
}
