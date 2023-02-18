package controllers

import (
	"database/sql"
	"fmt"
	"goblog/app/models/article"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	article, err := article.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		tpl, err := template.New("show.html").
			Funcs(template.FuncMap{
				"RouteName2URL":  route.Name2URL,
				"UInt64ToString": types.UInt64ToString,
			}).
			ParseFiles("resources/views/articles/show.html")
		logger.LogError(err)
		err = tpl.Execute(w, article)
		logger.LogError(err)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		tpl, err := template.ParseFiles("resources/views/articles/index.html")
		logger.LogError(err)

		err = tpl.Execute(w, articles)
		logger.LogError(err)
	}
}
