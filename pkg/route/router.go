package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Name2URL(routeName string, pairs ...string) string {
	var router *mux.Router
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	}

	return url.String()
}

func GetRouteVariable(name string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[name]
}
