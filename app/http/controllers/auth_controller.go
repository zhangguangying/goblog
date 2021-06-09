package controllers

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models/user"
	"github.com/zhangguangying/goblog/app/requests"
	"github.com/zhangguangying/goblog/pkg/auth"
	"github.com/zhangguangying/goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 表单验证
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"User":   _user,
			"Errors": errs,
		}, "auth.register")
	} else {
		_user.Create()
		if _user.ID > 0 {
			fmt.Fprintf(w, "创建用户成功，ID为：", _user.GetStringId())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 服务器内部错误")
		}
	}
}

// Login 显示登录表单
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 处理登录表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")
	if err := auth.Attempt(email, pass); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": pass,
		}, "auth.login")
	}
}
