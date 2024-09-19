package userLogging

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterUserLoggingApi(parentRoute iris.Party) *mvc.Application {

	container := parentRoute.ConfigureContainer().Container
	controller := new(controller.UserLoggingController)

	wrapper := mvc.New(parentRoute)

	var handleFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

		activator.Handle(
			"POST", "/login", "LogIn",
			middleware.BindPresenters[requestPresenter.LoginRequest, responsePresenter.LoginResponse](container),
		)
	}

	wrapper.Handle(
		controller,
		handleFunc,
	)

	return wrapper
}
