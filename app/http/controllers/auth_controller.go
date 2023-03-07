package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct{}

func (a AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (a AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	_user := user.User{
		Name:     name,
		Password: password,
		Email:    email,
	}
	_user.Create()
	if _user.ID > 0 {
		fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建用户失败，请联系管理员")
	}
}
