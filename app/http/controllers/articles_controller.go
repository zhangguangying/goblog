package controllers

import (
	"database/sql"
	"fmt"
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/app/policies"
	"github.com/zhangguangying/goblog/app/requests"
	"github.com/zhangguangying/goblog/pkg/auth"
	"github.com/zhangguangying/goblog/pkg/flash"
	"github.com/zhangguangying/goblog/pkg/logger"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/view"
	"gorm.io/gorm"
	"log"
	"net/http"
	"unicode/utf8"
)

type ArticlesController struct {
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
		if !policies.CanModifyArticle(_article) {
			flash.Warning("未授权操作")
			http.Redirect(w, r, "/", http.StatusFound)
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
		if !policies.CanModifyArticle(_article) {
			flash.Warning("未授权操作")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			errors := requests.ValidatorArticleForm(_article)

			if len(errors) == 0 {
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
				data := view.D{
					"Article": _article,
					"Errors":  errors,
				}
				view.Render(w, data, "articles.edit", "articles._form_field")
			}
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
		if !policies.CanModifyArticle(article) {
			flash.Warning("未授权操作")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			url := route.Name2URL("articles.update", "id", id)
			data := view.D{
				"Article": article,
				"URL":     url,
				"Errors":  view.D{},
			}
			view.Render(w, data, "articles.edit", "articles._form_field")
		}
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
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
	_article := article.Article{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		UserID: auth.User().ID,
	}

	errors := requests.ValidatorArticleForm(_article)
	fmt.Println(errors)
	if len(errors) == 0 {
		_article.Create()

		if _article.ID > 0 {
			flash.Success("创建成功")
			http.Redirect(w, r, route.Name2URL("articles.show", "id", _article.GetStringId()), http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		data := view.D{
			"Article": _article,
			"Errors":  errors,
		}
		view.Render(w, data, "articles.create", "articles._form_field")
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
		view.Render(w, view.D{"Article": _article}, "articles.show", "articles._article_meta")
	}
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, pageData, err := article.GetAll(r, 2)
	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 服务器内部错误")
	} else {
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pageData,
		}, "articles.index", "articles._article_meta")
	}
}
