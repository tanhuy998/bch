package authManipulationApi

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"
	"app/internal/middlewareHelper"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterRoleApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/roles")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthRoleManipulationController)

	wrapper := mvc.New(router)

	wrapper.Router.Use(
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
	)

	var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

		activator.Handle(
			"GET", "/", "GetAllRoles",
			middleware.BindPresenters[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse](container),
		)

		// activator.Handle(
		// 	"POST", "/", "CreateRole",
		// 	middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](container),
		// )

		activator.Handle(
			"POST", "/group/{groupUUID}/user/{userUUID}", "GrantCommandGroupRolesToUser",
			middleware.BindPresenters[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse](container),
		)
	}

	wrapper.Handle(
		controller,
		hanldFunc,
	)

	return wrapper
}
