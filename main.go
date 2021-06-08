package main

import (
	"github.com/zhangguangying/goblog/app/http/middlewares"
	"github.com/zhangguangying/goblog/bootstrap"
	"net/http"
)

func main() {
	bootstrap.SetDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
