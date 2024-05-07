package middleware

import (
	"app/app/model"

	"github.com/kataras/iris/v12"
)

func Authorization(roles []model.UserGroup) iris.Handler {

	return func(ctx iris.Context) {

		user := ctx.Values().Get(AUTH_USER)

		if user == nil {

		}

		ctx.Next()
	}
}
