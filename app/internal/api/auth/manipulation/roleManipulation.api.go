package authManipulationApi

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRoleApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/roles")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthCommandGroupManipulationController)

	wrapper := mvc.New(router)

	var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

		activator.Handle(
			"POST", "/", "CreateRole",
			middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](container),
		)

	}

	wrapper.Handle(
		controller,
		hanldFunc,
	)

	return wrapper
}