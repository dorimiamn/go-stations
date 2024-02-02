package middleware

import (
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler/ctxvalue"
	"github.com/mileusna/useragent"
)

func GetUserAgent(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user_agent := useragent.Parse(r.UserAgent())
		ctx := ctxvalue.SetUserAgent(r.Context(), &user_agent)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
