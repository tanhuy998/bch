package authReportApi

import (
	"app/infrastructure/http/api/v1/controller"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterAuthReportApi(parentRouter iris.Party) *mvc.Application {

	container := parentRouter.ConfigureContainer().Container

	parentRouter.Use(
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
	)

	controller := new(controller.AuthReportController)
	wrapper := mvc.New(parentRouter)

	wrapper.Handle(
		controller,
	)

	return wrapper
}
