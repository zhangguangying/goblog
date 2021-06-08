package controllers

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "文章不存在")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "服务器内部错误")
		}
	} else {
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
			"Int64ToString": types.Int64ToString,
		}).
			ParseFiles("resources/views/articles/show.gohtml")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, article)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 服务器内部错误")
	} else {
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		if err != nil {
			logger.LogError(err)
		}
		tmpl.Execute(w, articles)
	}

}
