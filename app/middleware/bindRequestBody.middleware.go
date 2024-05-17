package middleware

import "github.com/kataras/iris/v12"

func BindRequestBody[T any]() iris.Handler {

	return func(ctx iris.Context) {

		body := new(T)

		err := ctx.ReadJSON(body)

		if err != nil {

			ctx.StopWithJSON(400, body)
			return
		}

		ctx.RegisterDependency(body)
		ctx.Next()
	}
}
