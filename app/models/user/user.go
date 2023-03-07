package user

import (
	"goblog/app/models"
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

type User struct {
	models.BaseModel

	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);unique" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	ConfirmPassword string `gorm:"-" valid:"confirm_password"`
}

func (u User) Create() error {
	if err := model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
