package refreshTokenClientService

import (
	"app/internal/bootstrap"
	libError "app/internal/lib/error"
	refreshTokenServicePort "app/port/refreshToken"
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

	str := c.GetCookie(key)

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

	options := []irisContext.CookieOption{
		irisContext.CookiePath("/refresh"),
		irisContext.CookieHTTPOnly(true),
		irisContext.CookieSameSite(http.SameSiteStrictMode),
	}

	expire := refreshToken.GetExpireTime()

	if expire != nil {

		options = append(options, irisContext.CookieExpires(time.Until(*expire)))
	}

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		ops := append(options, irisContext.CookieDomain(hostname))

		c.SetCookieKV(
			key, rt, ops...,
		)
	}

	return nil
}

func (this *RefreshTokenClientService) Remove(ctx context.Context) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("RefreshTokenClientService error: invalid context given (not type of iris.Context)"))
	}

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		c.SetCookieKV(
			key, "",
			irisContext.CookiePath("/refresh"),
			irisContext.CookieHTTPOnly(true),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
			irisContext.CookieDomain(hostname),
			irisContext.CookieExpires(0),
		)
	}

	return nil
}
