package authSignaturesApi

import (
	"app/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterSignaturesApi(parentRouter router.Party) *mvc.Application {

	router := parentRouter.Party("/signatures")

	wrapper := mvc.New(router)

	wrapper.Handle(new(controller.AuthSignaturesController))

	return wrapper
}
