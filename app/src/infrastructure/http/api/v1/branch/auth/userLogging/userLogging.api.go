package userLogging

import (
	"app/src/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterUserLoggingApi(parentRoute iris.Party) *mvc.Application {

	//container := parentRoute.ConfigureContainer().Container
	controller := new(controller.UserLoggingController)

	wrapper := mvc.New(parentRoute)

	// var handleFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

	// 	activator.Handle(
	// 		"POST", "/login", "LogIn",
	// 		middleware.BindPresenters[requestPresenter.LoginRequest, responsePresenter.LoginResponse](container),
	// 	)

	// 	activator.Handle(
	// 		"POST", "/refresh", "Refresh",
	// 		middleware.BindPresenters[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse](container),
	// 	)
	// }

	wrapper.Handle(
		controller,
		//handleFunc,
	)

	return wrapper
}
