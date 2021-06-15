package article

import (
	"github.com/zhangguangying/goblog/app/models"
	"github.com/zhangguangying/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title string `valid:"title"`
	Body  string `valid:"body"`
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringId())
}
