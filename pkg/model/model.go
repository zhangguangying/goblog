package model

import (
	"goblog/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectionDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/goblog?charset=utf8&parseTime=true&loc=Local",
	})

	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})

	logger.LogError(err)

	return DB
}
