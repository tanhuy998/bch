package authManipulationApi

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func RegisterCommandGroupApi(parentRoute router.Party) *mvc.Application {

	router := parentRoute.Party("/groups")

	container := router.ConfigureContainer().Container
	controller := new(controller.AuthCommandGroupManipulationController)

	wrapper := mvc.New(router)

	var hanldFunc mvc.OptionFunc = func(activator *mvc.ControllerActivator) {

		activator.Handle(
			"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
			middleware.BindPresenters[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups](container),
		)

		// activator.Handle(
		// 	"GET", "/", "GetAllGroups",
		// )

		activator.Handle(
			"POST", "/{groupUUID}/grant/user/{userUUID}", "GrantCommandGroupRolesToUser",
			middleware.BindPresenters[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse](container),
		)

		activator.Handle(
			"POST", "/", "CreateGroup",
			middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](container),
		)

		activator.Handle(
			"POST", "/{groupUUID:uuid}/user/{userUUID:uuid}", "AddUserToGroup",
			middleware.BindPresenters[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse](container),
		)
	}

	wrapper.Handle(
		controller,
		hanldFunc,
	)

	return wrapper
}
