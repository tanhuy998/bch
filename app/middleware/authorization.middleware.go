package middleware

import (
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
)

type AuthorizationLicense struct {
	Fields []authService.AuthorizationField
	Groups []authService.AuthorizationGroup
	Claims []authService.AuthorizationClaim
}

func Authorization(option AuthorizationLicense) iris.Handler {

	return func(ctx iris.Context) {

		user := ctx.Values().Get(AUTH_USER)

		if user == nil {

		}

		ctx.Next()
	}
}

func AuthorizeGroups() {

}

func AuthorizeClaims() {

}
