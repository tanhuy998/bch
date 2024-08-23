package authManipulationApi

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterUserApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/users")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthUserManipulationController)

	wrapper := mvc.New(router)

	var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

		activator.Handle(
			"POST", "/", "CreateUser",
			middleware.BindPresenters[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter](container),
		)

		activator.Handle(
			"GET", "/group/{groupUUID:uuid}", "GetGroupUsers",
			middleware.BindPresenters[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse](container),
		)
	}

	wrapper.Handle(
		controller,
		hanldFunc,
	)

	return wrapper
}
