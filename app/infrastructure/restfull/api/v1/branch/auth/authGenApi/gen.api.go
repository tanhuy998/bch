package authGenApi

import (
	"app/infrastructure/http/api/v1/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterGenAPI(parentRouter router.Party) *mvc.Application {

	router := parentRouter.Party("/gen")

	wrapper := mvc.New(router)

	wrapper.Handle(new(controller.AuthGeneralController))

	return wrapper
}

func API(parentRouter router.Party) {

	launcher := common.NewAPILauncher(parentRouter)

	launcher.LaunchAPIOf(
		new(AuthGeneralController),
	)
}
