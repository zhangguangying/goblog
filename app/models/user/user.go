package user

import (
	"goblog/app/models"
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	Email    string `gorm:"column:email;type:varchar(255);default null;unique"`
	Password string `gorm:"column:password;type:varchar(255)"`
}

func (u User) Create() error {
	if err := model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
