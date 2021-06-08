package controllers

import (
	"database/sql"
	"fmt"
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"unicode/utf8"
)

type ArticlesController struct {
}

type ArticlesStoreData struct {
	Title, Body string
	URL         string
	Errors      map[string]string
}

func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		rowsAffect, err := _article.Delete()
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		} else {
			if rowsAffect > 0 {
				indexUrl := route.Name2URL("articles.index")
				http.Redirect(w, r, indexUrl, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "404 文章未找到")
			}
		}
	}
}

func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 文章未找到！")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		}
	} else {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")
		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {
			_article.Title = title
			_article.Body = body

			rs, err := _article.Update()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			}
			if rs > 0 {
				showUrl := route.Name2URL("articles.show", "id", id)
				http.Redirect(w, r, showUrl, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改")
			}
		} else {
			url := route.Name2URL("articles.update", "id", id)
			data := ArticlesStoreData{
				Title:  _article.Title,
				Body:   _article.Body,
				URL:    url,
				Errors: errors,
			}
			tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
			logger.LogError(err)
			tmpl.Execute(w, data)
		}
	}
}

func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	article, err := article.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "文章不存在")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			fmt.Fprintf(w, "服务器内部错误")
		}
	} else {
		url := route.Name2URL("articles.update", "id", id)
		data := ArticlesStoreData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    url,
			Errors: nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, data)
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeUrl := route.Name2URL("articles.store")
	data := ArticlesStoreData{
		Title:  "",
		Body:   "",
		URL:    storeUrl,
		Errors: nil,
	}

	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		fmt.Fprint(w, err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprint(w, err)
	}
}

func validateArticleFormData(title, body string) map[string]string {
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于等于 10 个字节"
	}

	return errors
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
		_article.Create()

		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+strconv.FormatInt(int64(_article.ID), 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		storeUrl := route.Name2URL("articles.store")

		data := ArticlesStoreData{
			Title:  title,
			Body:   body,
			URL:    storeUrl,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, data)
	}
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "文章不存在")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "服务器内部错误")
		}
	} else {
		viewDir := "resources/views"
		files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
		logger.LogError(err)
		files = append(files, viewDir+"/articles/show.gohtml")
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
			"Int64ToString": types.Int64ToString,
		}).
			ParseFiles(files...)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "app", _article)
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 服务器内部错误")
	} else {
		viewDir := "resources/views"
		files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
		logger.LogError(err)
		files = append(files, viewDir+"/articles/index.gohtml")
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			logger.LogError(err)
		}
		tmpl.ExecuteTemplate(w, "app", articles)
	}

}
