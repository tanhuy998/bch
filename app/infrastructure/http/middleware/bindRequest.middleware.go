package middleware

import (
	"app/infrastructure/http/middleware/middlewareHelper"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type ()

func BindRequest[Request_T any](
	container *hero.Container,
	initializer ...middlewareHelper.PresenterInitializer[Request_T, EmptyPresenter],
) iris.Handler {

	return BindPresenters(container, initializer...)
}
