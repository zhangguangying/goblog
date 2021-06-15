package policies

import (
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
