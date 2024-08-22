package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthRoleManipulationController struct {
		GetAllRolesUseCase usecase.IGetAllRoles
	}
)

func (this *AuthRoleManipulationController) CreateRole() {

}

func (this *AuthRoleManipulationController) GetAllRoles(
	input *requestPresenter.GetAllRolesRequest,
	output *responsePresenter.GetAllRolesResponse,
) (mvc.Result, error) {

	return this.GetAllRolesUseCase.Execute(input, output)
}
