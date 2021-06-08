package controllers

import (
	"database/sql"
	"fmt"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/types"
	"html/template"
	"log"
	"net/http"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := getArticleById(id)
	if err != nil {
		if err == sql.ErrNoRows {
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
f