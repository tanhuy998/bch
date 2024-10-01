package controller

import (
	"app/src/infrastructure/http/middleware"
	"app/src/infrastructure/http/middleware/middlewareHelper"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	usecase "app/src/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthCommandGroupManipulationController struct {
		CreateCommandGroupUseCase          usecase.ICreateCommandGroup
		AddUserToCommandGroupUseCase       usecase.IAddUserToCommandGroup
		GetParitcipatedCommandGroupUseCase usecase.IGetParticipatedCommandGroups
	}
)

func (this *AuthCommandGroupManipulationController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"GET", "/participated/user/{userUUID:uuid}", "GetParticipatedGroups",
		middleware.BindPresenters[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	// activator.Handle(
	// 	"GET", "/", "GetAllGroups",
	// )

	activator.Handle(
		"POST", "/", "CreateGroup",
		middleware.BindPresenters[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)

	activator.Handle(
		"POST", "/{groupUUID:uuid}/user/{userUUID:uuid}", "AddUserToGroup",
		middleware.BindPresenters[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse](
			container,
			middlewareHelper.UseAuthority,
		),
	)
}

func (this *AuthCommandGroupManipulationController) CreateGroup(
	input *requestPresenter.CreateCommandGroupRequest,
	output *responsePresenter.CreateCommandGroupResponse,
) (mvc.Result, error) {

	return this.CreateCommandGroupUseCase.Execute(input, output)
}

func (this *AuthCommandGroupManipulationController) AddUserToGroup(
	input *requestPresenter.AddUserToCommandGroupRequest,
	output *responsePresenter.AddUserToCommandGroupResponse,
) (mvc.Result, error) {

	return this.AddUserToCommandGroupUseCase.Execute(input, output)
}

func (this *AuthCommandGroupManipulationController) GetAllGroups() {

}

func (this *AuthCommandGroupManipulationController) GetParticipatedGroups(
	input *requestPresenter.GetParticipatedGroups,
	output *responsePresenter.GetParticipatedGroups,
) (mvc.Result, error) {

	return this.GetParitcipatedCommandGroupUseCase.Execute(input, output)
}

// func (this *AuthCommandGroupManipulationController) GrantCommandGroupRolesToUser(
// 	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
// 	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
// ) (mvc.Result, error) {

// 	return this.GrantCommandGroupRolesToUserUseCase.Execute(input, output)
// }
