package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetAllRoles interface {
		Execute(
			input *requestPresenter.GetAllRolesRequest,
			output *responsePresenter.GetAllRolesResponse,
		) (mvc.Result, error)
	}

	GetAllRolesUseCase struct {
		GetAllRolesService authService.IGetAllRoles
		ActionResult       actionResultService.IActionResult
	}
)

func (this *GetAllRolesUseCase) Execute(
	input *requestPresenter.GetAllRolesRequest,
	output *responsePresenter.GetAllRolesResponse,
) (mvc.Result, error) {

	ret, err := this.GetAllRolesService.Serve()

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = ret

	return this.ActionResult.ServeResponse(output)
}
