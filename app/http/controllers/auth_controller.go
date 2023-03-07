package controllers

import (
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct{}

func (a AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (a AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

}
