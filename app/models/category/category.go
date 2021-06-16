package category

import (
	"github.com/zhangguangying/goblog/app/models"
	"github.com/zhangguangying/goblog/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

func (c Category) Link() string {
	return route.Name2URL("categories.show", "id", c.GetStringId())
}
