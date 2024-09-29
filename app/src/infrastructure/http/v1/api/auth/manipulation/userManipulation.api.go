package authManipulationApi

import (
	"app/internal/controller"
	"app/internal/middleware"
	"app/internal/middlewareHelper"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterUserApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/users")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthUserManipulationController)

	wrapper := mvc.New(router)

	wrapper.Router.Use(
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
	)

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
