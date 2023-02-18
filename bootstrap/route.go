package bootstrap

import (
	"goblog/routes"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	routes.RegisteWebRoutes(router)
	return router
}
