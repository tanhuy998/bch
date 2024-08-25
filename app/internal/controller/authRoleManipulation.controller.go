package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type (
	AuthRoleManipulationController struct {
		GetAllRolesUseCase                  usecase.IGetAllRoles
		GrantCommandGroupRolesToUserUseCase usecase.IGrantCommandGroupRolesToUser
	}
)

func (this *AuthRoleManipulationController) GrantCommandGroupRolesToUser(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
) (mvc.Result, error) {

	return this.GrantCommandGroupRolesToUserUseCase.Execute(input, output)
}

func (this *AuthRoleManipulationController) GetAllRoles(
	input *requestPresenter.GetAllRolesRequest,
	output *responsePresenter.GetAllRolesResponse,
) (mvc.Result, error) {

	return this.GetAllRolesUseCase.Execute(input, output)
}
