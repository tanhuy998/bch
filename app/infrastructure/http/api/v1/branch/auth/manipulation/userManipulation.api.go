package authManipulationApi

import (
	"app/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterUserApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/users")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthUserManipulationController).BindDependencies(container)

	wrapper := mvc.New(router)

	// wrapper.Router.Use(
	// 	middleware.Auth(
	// 		container,
	// 	),
	// )

	// var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

	// 	activator.Handle(
	// 		"POST", "/", "CreateUser",
	// 		middleware.BindPresenters[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter](
	// 			container,
	// 			middlewareHelper.UseAuthority,
	// 		),
	// 	)

	// 	activator.Handle(
	// 		"GET", "/group/{groupUUID:uuid}", "GetGroupUsers",
	// 		middleware.BindPresenters[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse](
	// 			container,
	// 			middlewareHelper.UseAuthority,
	// 		),
	// 	)

	// 	activator.Handle(
	// 		"PATCH", "/{userUUID:uuid}", "ModifyUser",
	// 		middleware.BindPresenters[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse](
	// 			container,
	// 			middlewareHelper.UseAuthority,
	// 		),
	// 	)
	// }

	wrapper.Handle(
		controller,
		//hanldFunc,
	)

	return wrapper
}
