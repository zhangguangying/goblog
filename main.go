package main

import (
	"database/sql"
	"goblog/app/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var db *sql.DB
var router *mux.Router

func main() {
	bootstrap.SetupDB()

	router = bootstrap.SetupRouter()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
