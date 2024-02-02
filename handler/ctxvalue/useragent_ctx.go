package ctxvalue

import (
	"context"

	"github.com/mileusna/useragent"
)

type contextKey struct {}

var userAgentKey = contextKey{}

func GetUserAgent(ctx context.Context) (*useragent.UserAgent, bool) {
	userAgent, ok := ctx.Value(userAgentKey).(*useragent.UserAgent)
	return userAgent, ok
}

func SetUserAgent(ctx context.Context, u *useragent.UserAgent) context.Context {
	return context.WithValue(ctx, userAgentKey, u)
}
