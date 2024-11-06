package refreshTokenClientService

import (
	"app/internal/bootstrap"
	libError "app/internal/lib/error"
	refreshTokenServicePort "app/port/refreshToken"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"context"

	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
)

const (
	key = "refresh-token"
)

type (
	RefreshTokenClientService struct {
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *RefreshTokenClientService) Read(ctx context.Context) (refreshTokenServicePort.IRefreshToken, error) {

	c, ok := ctx.(iris.Context)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("RefreshTokenClientService error: invalid context given (not type of iris.Context)"))
	}

	b, _ := json.Marshal(c.Request().Header)

	fmt.Println(string(b))

	str := c.GetCookie(key)
	fmt.Println(c.Request().Cookies())
	if str == "" {

		return nil, nil
	}

	rt, err := this.RefreshTokenManipulator.Read(str)

	if err != nil {

		return nil, err
	}

	return rt, nil
}

func (this *RefreshTokenClientService) Write(ctx context.Context, refreshToken refreshTokenServicePort.IRefreshToken) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("RefreshTokenClientService error: invalid context given (not type of iris.Context)"))
	}

	rt, err := this.RefreshTokenManipulator.SignString(refreshToken)

	if err != nil {

		return err
	}

	c.AddCookieOptions(
		irisContext.CookieHTTPOnly(true),
		irisContext.CookieSameSite(http.SameSiteStrictMode),
	)

	expire := refreshToken.GetExpireTime()

	if expire != nil {

		c.AddCookieOptions(
			irisContext.CookieExpires(time.Until(*expire)),
		)
	}

	for _, domain := range bootstrap.GetDomainNames() {

		if domain == "*" {

			continue
		}

		c.SetCookieKV(
			key, rt,
			irisContext.CookieDomain(domain),
			//irisContext.CookiePath("/auth/refresh"),
			irisContext.CookiePath("/auth/signatures"),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
		)

		// c.SetCookieKV(
		// 	key, rt,
		// 	irisContext.CookieDomain(domain),
		// 	//irisContext.CookiePath("/tenants/switch"),
		// 	irisContext.CookiePath("/auth/signatures/tenant/"),
		// )

		// c.SetCookieKV(
		// 	key, rt,
		// 	irisContext.CookieDomain(domain),
		// 	irisContext.CookiePath("/auth/logout"),
		// )

	}

	return nil
}

func (this *RefreshTokenClientService) Remove(ctx context.Context) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("RefreshTokenClientService error: invalid context given (not type of iris.Context)"))
	}

	c.AddCookieOptions(
		irisContext.CookieHTTPOnly(true),
		irisContext.CookieSameSite(http.SameSiteStrictMode),
	)

	for _, domain := range bootstrap.GetDomainNames() {

		if domain == "*" {

			continue
		}

		c.RemoveCookie(
			key,
			//irisContext.CookiePath("/auth/refresh"),
			irisContext.CookiePath("/auth/signatures"),
			irisContext.CookieDomain(domain),
		)

		// c.RemoveCookie(
		// 	key,
		// 	irisContext.CookiePath("/tenants/switch"),
		// 	irisContext.CookieDomain(domain),
		// )

		// c.RemoveCookie(
		// 	key,
		// 	irisContext.CookiePath("/auth/logout"),
		// 	irisContext.CookieDomain(domain),
		// )
	}

	return nil
}
