package bootstrap

import (
	"goblog/pkg/route"
	"goblog/routes"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	routes.RegisteWebRoutes(router)
	route.SetupRouter(router)
	return router
}
