package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func CheckAuth(r *http.Request) bool {
	// 認証情報を取得
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	// 認証情報を照合
	return username == os.Getenv("BASIC_AUTH_USER_ID") && password == os.Getenv("BASIC_AUTH_PASSWORD")
}

func BasicAuth(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !CheckAuth(r) {
			fmt.Println("Unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("Authorized")
	}
	return http.HandlerFunc(fn)
}