package user

import (
	"github.com/zhangguangying/goblog/pkg/password"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = password.Hash(u.Password)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}
