package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

func GetAgent(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user_agent := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), "user_agent", user_agent)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}