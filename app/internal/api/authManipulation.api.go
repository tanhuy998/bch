package api

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initAuthApi(app *iris.Application) *mvc.Application {

	router := app.Party("/auth/man")

	container := router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	}).Container

	controller := new(controller.AuthManipulationController)

	wrapper := mvc.New(router)

	wrapper.Handle(
		controller,
		applyRoutes(func(activator *mvc.ControllerActivator) {

			activator.Handle(
				"POST", "/user", "CreateUser",
				middleware.BindPresenters[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter](container),
			)
		}),
	)

	return wrapper
}
