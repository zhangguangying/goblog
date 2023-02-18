package routes

import (
	"goblog/app/http/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisteWebRoutes(r *mux.Router) {
	pc := new(controller.PagesController)

	r.HandleFunc("/", pc.Home).Name("home")
	r.HandleFunc("/about", pc.About).Name("about")
	// 404页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
}
