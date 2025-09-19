package authReportApi

import (
	"app/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterAuthReportApi(parentRouter iris.Party) *mvc.Application {

	//container := parentRouter.ConfigureContainer().Container

	controller := new(controller.AuthReportController)
	wrapper := mvc.New(parentRouter)

	wrapper.Handle(
		controller,
	)

	return wrapper
}
