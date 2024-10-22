package generalTokenClientService

import (
	"app/internal/bootstrap"
	libError "app/internal/lib/error"
	generalTokenServicePort "app/port/generalToken"

	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	irisContext "github.com/kataras/iris/v12/context"
)

const (
	GENERAL_TOKEN = "general-token"
)

type (
	GeneralTokenClientService struct {
		GeneralTokenManipulator generalTokenServicePort.IGeneralTokenManipulator
	}
)

func (this *GeneralTokenClientService) Read(ctx context.Context) (generalTokenServicePort.IGeneralToken, error) {

	c, ok := ctx.(iris.Context)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("GeneralTokenClientService error: invalid context given (not type of iris.Context)"))
	}

	str := c.GetCookie(GENERAL_TOKEN)

	if str == "" {

		return nil, nil
	}

	ret, err := this.GeneralTokenManipulator.Read(str)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GeneralTokenClientService) Write(ctx context.Context, generalToken generalTokenServicePort.IGeneralToken) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("GeneralTokenClientService error: invalid token given (not type of iris.Context)"))
	}

	gt, err := this.GeneralTokenManipulator.SignString(generalToken)

	if err != nil {

		return err
	}

	options := []irisContext.CookieOption{
		irisContext.CookieHTTPOnly(true),
		irisContext.CookieSameSite(http.SameSiteStrictMode),
	}

	expire := generalToken.GetExpiretime()

	if expire != nil {

		options = append(options, irisContext.CookieExpires(time.Until(*expire)))
	}

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		c.SetCookieKV(
			GENERAL_TOKEN,
			gt,
			append(
				options,
				irisContext.CookieDomain(hostname),
				irisContext.CookiePath("/tenants/switch"),
			)...,
		)

		c.SetCookieKV(
			GENERAL_TOKEN,
			gt,
			append(
				options,
				irisContext.CookieDomain(hostname),
				irisContext.CookiePath("/auth/nav"),
			)...,
		)
	}

	return nil
}

func (this *GeneralTokenClientService) Remove(ctx context.Context) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("GeneralTokenClientService error: invalid token given (not type of iris.Context)"))
	}

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		c.RemoveCookie(
			GENERAL_TOKEN,
			irisContext.CookiePath("/tenants/switch"),
			irisContext.CookieHTTPOnly(true),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
			irisContext.CookieDomain(hostname),
			irisContext.CookieExpires(0),
		)

		c.RemoveCookie(
			GENERAL_TOKEN,
			irisContext.CookiePath("/auth/nav"),
			irisContext.CookieHTTPOnly(true),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
			irisContext.CookieDomain(hostname),
			irisContext.CookieExpires(0),
		)
	}

	return nil
}
