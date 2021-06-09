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
