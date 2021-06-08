package model

import (
	"github.com/zhangguangying/goblog/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gologger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	config := mysql.New(mysql.Config{
		DSN: "homestead:secret@tcp(192.168.10.10:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gologger.Default.LogMode(gologger.Warn),
	})
	logger.LogError(err)
	return DB
}
