package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
	"github.com/TechBowl-japan/go-stations/handler/ctx/contextValue"
)

func GetUserAgent(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user_agent := useragent.Parse(r.UserAgent())

		ctx := context.WithValue(r.Context(), userAgentKey, user_agent)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
