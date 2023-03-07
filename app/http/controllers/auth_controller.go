package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct{}

func (a AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (a AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		ConfirmPassword: r.PostFormValue("confirm_password"),
	}
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// data, _ := json.MarshalIndent(errs, "", " ")
		// fmt.Fprint(w, string(data))
		fmt.Println(errs)
		view.RenderSimple(w, &view.D{
			"Errors": errs,
		}, "auth.register")
	} else {
		_user.Create()
		if _user.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败，请联系管理员")
		}
	}

}
