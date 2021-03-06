package controllers

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/app/models/user"
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
}

func (*UserController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_user, err := user.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 用户未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		}
	} else {
		articles, err := article.GetByUserId(_user.GetStringId())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}
}
