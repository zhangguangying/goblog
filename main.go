package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zhangguangying/goblog/bootstrap"
	"github.com/zhangguangying/goblog/pkg/database"
	"github.com/zhangguangying/goblog/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var router *mux.Router
var db *sql.DB

type Article struct {
	Title, Body string
	ID          int64
}

func (a Article) Delete() (rowsAffect int64, err error) {
	rs, err := db.Exec("delete from articles where id = " + strconv.FormatInt(a.ID, 10))
	if err != nil {
		return 0, err
	}
	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}
	return 0, nil
}

type ArticlesStoreData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

func forceHTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		h.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func getArticleById(id string) (Article, error) {
	var article Article
	query := "select * from articles where id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetDB()
	router = bootstrap.SetupRoute()

	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).Methods("POST").Name("articles.delete")

	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}

func articlesDeleteHandler(writer http.ResponseWriter, request *http.Request) {
	id := getRouteVariable("id", request)
	article, err := getArticleById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(writer, "404 文章未找到")
		} else {
			logger.LogError(err)
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(writer, "500 服务器内部错误")
		}
	} else {
		rowsAffect, err := article.Delete()
		if err != nil {
			logger.LogError(err)
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "500 服务器内部错误")
		} else {
			if rowsAffect > 0 {
				indexUrl, _ := router.Get("articles.index").URL()
				http.Redirect(writer, request, indexUrl.String(), http.StatusFound)
			} else {
				writer.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(writer, "404 文章未找到")
			}
		}
	}
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
