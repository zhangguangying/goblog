package route

import (
	"github.com/gorilla/mux"
	"github.com/zhangguangying/goblog/pkg/routes"
	"net/http"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameter string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameter]
}
