package middleware

import (
	authService "app/app/service/auth"
	"fmt"

	"github.com/kataras/iris/v12"
)

func Authorize(licenses ...authService.AuthorizationLicense) iris.Handler {

	return func(ctx iris.Context) {

		// user := ctx.Values().Get(AUTH_USER)

		// if user == nil {

		// }
		fmt.Println("2")
		ctx.Next()
	}
}

func AuthorizeGroups() {

}

func AuthorizeClaims() {

}
