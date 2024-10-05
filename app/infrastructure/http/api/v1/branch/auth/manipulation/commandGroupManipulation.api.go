package authManipulationApi

import (
	"app/infrastructure/http/api/v1/controller"
	"app/infrastructure/http/middleware"
	"app/infrastructure/http/middleware/middlewareHelper"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterCommandGroupApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/groups")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthCommandGroupManipulationController).BindDependencies(container)

	wrapper := mvc.New(router)
	wrapper.Router.Use(
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
	)

	// var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

	// 	activator.Handle(
	// 		"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
	// 		middleware.BindPresenters[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups](
	// 			container,
	// 			middlewareHelper.UseAuthority,
	// 		),
	// 	)

	// 	// activator.Handle(
	// 	// 	"GET", "/", "GetAllGroups",
	// 	// )

	// 	activator.Handle(
	// 		"POST", "/", "CreateGroup",
	// 		middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](
	// 			container,
	// 			middlewareHelper.UseAuthority,
	// 		),
	// 	)

	// 	activator.Handle(
	// 		"POST", "/{groupUUID:uuid}/user/{userUUID:uuid}", "AddUserToGroup",
	// 		middleware.BindPresenters[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse](
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
