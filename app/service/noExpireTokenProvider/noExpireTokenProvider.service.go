package noExpireTokenProvider

import "context"

const (
	key_no_expire = "NO_EXPIRE_TOKEN"
)

type (
	NoExpireTokenProvider struct {
	}
)

func (this *NoExpireTokenProvider) CtxNoExpireKey() string {

	return key_no_expire
}

func (this *NoExpireTokenProvider) IsNoExpire(ctx context.Context) bool {

	v, ok := ctx.Value(key_no_expire).(bool)

	return ok && v
}
