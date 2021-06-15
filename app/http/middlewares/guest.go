package middlewares

import (
	"github.com/zhangguangying/goblog/pkg/auth"
	"github.com/zhangguangying/goblog/pkg/flash"
	"net/http"
)

func Guest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {
			flash.Warning("登录用户无法访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
