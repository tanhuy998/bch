package middleware

import (
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
)

func Authorization(licenses ...authService.AuthorizationLicense) iris.Handler {

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
