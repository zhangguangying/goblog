package user

import (
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/model"
)

func (user *User) Create() (err error) {
	if err := model.DB.Create(user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func Get(id string) (User, error) {
	var _user User
	if err := model.DB.Where("id = ?", id).First(&_user).Error; err != nil {
		logger.LogError(err)
		return _user, err
	}
	return _user, nil
}

func GetByEmail(email string) (User, error) {
	var _user User
	if err := model.DB.Where("email = ?", email).First(&_user).Error; err != nil {
		logger.LogError(err)
		return _user, err
	}
	return _user, nil
}
