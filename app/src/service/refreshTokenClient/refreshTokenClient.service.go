package refreshTokenClientService

import (
	"app/src/internal/bootstrap"
	refreshTokenServicePort "app/src/port/refreshToken"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

const (
	key = "refresh-token"
)

type (
	RefreshTokenClientService struct {
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *RefreshTokenClientService) Read(ctx iris.Context) string {

	return ctx.GetCookie(key)
}

func (this *RefreshTokenClientService) Write(ctx iris.Context, refreshToken string) error {

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		ctx.SetCookieKV(
			key,
			refreshToken,
			context.CookieExpires(this.RefreshTokenManipulator.DefaultExpireDuration()),
			context.CookieDomain(hostname),
			context.CookiePath("/refresh"),
			context.CookieHTTPOnly(true),
			context.CookieSameSite(http.SameSiteStrictMode),
		)
	}

	return nil
}
