package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthCommandGroupManipulationController struct {
		CreateCommandGroupUseCase    usecase.ICreateCommandGroup
		AddUserToCommandGroupUseCase usecase.IAddUserToCommandGroup
	}
)

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

func (this *AuthCommandGroupManipulationController) GetParticipatedGroups() {

}
