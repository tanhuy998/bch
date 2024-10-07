package noExpireTokenServicePort

import "context"

type (
	INoExpireTokenProvider interface {
		CtxNoExpireKey() string
		IsNoExpire(ctx context.Context) bool
	}
)
