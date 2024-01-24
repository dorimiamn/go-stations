package contextValue

import (
	"context"

	"github.com/mileusna/useragent"
)

type contextKey struct {}

var userAgentKey = contextKey{}

func UserAgent(ctx context.Context) (*useragent.UserAgent, bool) {
	useragent, ok := ctx.Value(userAgentKey).(*useragent.UserAgent)
	return useragent, ok
}

func SetUserAgent(ctx context.Context, u *useragent.UserAgent) context.Context {
	return context.WithValue(ctx, userAgentKey, u)
}
