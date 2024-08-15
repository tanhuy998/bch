package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthCommandGroupManipulationController struct {
		CreateCommandGroupUseCase usecase.ICreateCommandGroup
	}
)

func (this *AuthCommandGroupManipulationController) CreateGroup(
	input *requestPresenter.CreateCommandGroupRequest,
	output *responsePresenter.CreateCommandGroupResponse,
) (mvc.Result, error) {

	return this.CreateCommandGroupUseCase.Execute(input, output)
}
