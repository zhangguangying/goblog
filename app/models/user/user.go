package user

import "github.com/zhangguangying/goblog/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	Email    string `gorm:"column:email;type:varchar(255);default null;uniqueIndex"`
	Password string `gorm:"column:password;type:varchar(255)"`
}
