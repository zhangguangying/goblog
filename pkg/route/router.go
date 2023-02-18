package route

import (
	"goblog/routes"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
	routes.RegisteWebRoutes(Router)
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		// checkError(err)
		return ""
	}

	return url.String()
}

func GetRouteVariable(name string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[name]
}
