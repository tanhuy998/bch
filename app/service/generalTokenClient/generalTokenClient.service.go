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

	c.AddCookieOptions(
		irisContext.CookieHTTPOnly(true),
		irisContext.CookieSameSite(http.SameSiteStrictMode),
	)

	expire := generalToken.GetExpiretime()

	if expire != nil {

		c.AddCookieOptions(irisContext.CookieExpires(time.Until(*expire)))
	}

	for _, domain := range bootstrap.GetDomainNames() {

		if domain == "*" {

			continue
		}

		c.SetCookieKV(
			GENERAL_TOKEN,
			gt,
			irisContext.CookieDomain(domain),
			//irisContext.CookiePath("/tenants/switch"),
			irisContext.CookiePath("/auth/gen"),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
		)

		c.SetCookieKV(
			GENERAL_TOKEN,
			gt,
			irisContext.CookieDomain(domain),
			//irisContext.CookiePath("/tenants/switch"),
			irisContext.CookiePath("/auth/signatures/tenant"),
			irisContext.CookieSameSite(http.SameSiteStrictMode),
		)

		// c.SetCookieKV(
		// 	GENERAL_TOKEN,
		// 	gt,
		// 	irisContext.CookieDomain(domain),
		// 	irisContext.CookiePath("/auth/nav"),
		// )

		// c.SetCookieKV(
		// 	GENERAL_TOKEN,
		// 	gt,
		// 	irisContext.CookieDomain(domain),
		// 	irisContext.CookiePath("/auth/login"),
		// )
	}

	return nil
}

func (this *GeneralTokenClientService) Remove(ctx context.Context) error {

	c, ok := ctx.(iris.Context)

	if !ok {

		return libError.NewInternal(fmt.Errorf("GeneralTokenClientService error: invalid token given (not type of iris.Context)"))
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
			GENERAL_TOKEN,
			//irisContext.CookiePath("/tenants/switch"),
			irisContext.CookiePath("/auth/gen"),
			irisContext.CookieDomain(domain),
		)

		// c.RemoveCookie(
		// 	GENERAL_TOKEN,
		// 	irisContext.CookiePath("/auth/nav"),
		// 	irisContext.CookieDomain(domain),
		// )

		// c.RemoveCookie(
		// 	GENERAL_TOKEN,
		// 	irisContext.CookiePath("/auth/login"),
		// 	irisContext.CookieDomain(domain),
		// )
	}

	return nil
}
