package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Hello, 这里是 goblog</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系。</p>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:hyuiing@163.com\">hyuiing@163.com</a>")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.SplitN(r.URL.Path, "/", 3)[2]
		fmt.Fprint(w, "文章 ID:"+id)
	})
	router.HandleFunc("/articles", func(writer http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprint(writer, "文章列表")
		case http.MethodPost:
			fmt.Fprint(writer, "创建新文章")
		}
	})

	http.ListenAndServe(":3000", router)
}
