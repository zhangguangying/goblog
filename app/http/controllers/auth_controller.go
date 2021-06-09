package controllers

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models/user"
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

	// 验证通过插入用户，跳转首页
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	_user.Create()
	if _user.ID > 0 {
		fmt.Fprintf(w, "创建用户成功，ID为：", _user.GetStringId())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 服务器内部错误")
	}

	// 验证失败
}
