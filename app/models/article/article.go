package article

import (
	"github.com/zhangguangying/goblog/pkg/route"
	"strconv"
)

type Article struct {
	ID    int64
	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
