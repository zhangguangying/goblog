package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")

	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	fmt.Fprint(w, "访问文章列表")
}
func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	fmt.Fprint(w, "创建一篇文章")
}
func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID:"+id)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Name("home")
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	homeUrl, _ := router.Get("home").URL()
	fmt.Println("homeUrl:", homeUrl)
	articleUrl, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleUrl:", articleUrl)

	http.ListenAndServe(":3000", router)
}
