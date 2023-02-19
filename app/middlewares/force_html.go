package middlewares

import "net/http"

func ForceHTML(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		handler.ServeHTTP(w, r)
	})
}
