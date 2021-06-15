package article

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models"
	"github.com/zhangguangying/goblog/app/models/user"
	"github.com/zhangguangying/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title  string `gorm:"type:varchar(255)" valid:"title"`
	Body   string `gorm:"type:text" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringId())
}

func (a Article) CreatedAtDate() string {
	str := a.CreatedAt.Format("2006-01-02")
	fmt.Println(str)
	return str
}
