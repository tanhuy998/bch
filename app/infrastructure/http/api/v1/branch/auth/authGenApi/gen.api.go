package authGenApi

import (
	"app/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterGenAPI(parentRouter router.Party) *mvc.Application {

	router := parentRouter.Party("/gen")

	wrapper := mvc.New(router)

	wrapper.Handle(new(controller.AuthGeneralController))

	return wrapper
}
