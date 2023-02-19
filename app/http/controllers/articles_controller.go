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
	"strconv"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ArticlesController struct {
}

type ArticleFormData struct {
	Body, Title string
	URL         string
	Errors      map[string]string
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

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	postUrl := route.Name2URL("articles.store")

	data := ArticleFormData{
		Title:  "",
		Body:   "",
		URL:    postUrl,
		Errors: nil,
	}

	tpl, err := template.ParseFiles("resources/views/articles/create.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		err := _article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功, ID为"+strconv.FormatUint(_article.ID, 10))
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		storeUrl := route.Name2URL("articles.store")

		data := ArticleFormData{
			Title:  title,
			Body:   body,
			URL:    storeUrl,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.html")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}
}

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 记录不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		tpl, err := template.ParseFiles("resources/views/articles/edit.html")
		logger.LogError(err)

		url := route.Name2URL("articles.update", "id", id)
		data := ArticleFormData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    url,
			Errors: nil,
		}
		err = tpl.Execute(w, data)
		logger.LogError(err)
	}
}

func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 记录不存在")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {
			_article.Title = title
			_article.Body = body
			rowsAffected, err := _article.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			}
			if rowsAffected > 0 {
				showUrl := route.Name2URL("articles.show", "id", id)
				http.Redirect(w, r, showUrl, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {
			updateUrl := route.Name2URL("articles.update", "id", id)

			data := ArticleFormData{
				Title:  title,
				Body:   body,
				URL:    updateUrl,
				Errors: errors,
			}
			tmpl, err := template.ParseFiles("resources/views/articles/edit.html")
			if err != nil {
				panic(err)
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				panic(err)
			}
		}
	}
}

func validateArticleFormData(title, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度介于3-40"
	}

	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}
	return errors
}
