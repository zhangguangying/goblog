package user

import (
	"github.com/zhangguangying/goblog/app/models"
	"github.com/zhangguangying/goblog/pkg/password"
)

type User struct {
	models.BaseModel

	Name            string `gorm:"column:name;type:varchar(255);not null;uniqueIndex" valid:"name"`
	Email           string `gorm:"column:email;type:varchar(255);default null;uniqueIndex" valid:"email"`
	Password        string `gorm:"column:password;type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (u *User) ComparePassword(pass string) bool {
	return password.CheckHash(u.Password, pass)
}

func (u *User) Link() string {
	return ""
}
