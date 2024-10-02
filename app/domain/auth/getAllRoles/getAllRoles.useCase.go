package getAllRoleDomain

import (
	actionResultServicePort "app/port/actionResult"
	authServicePort "app/port/auth"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

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
		GetAllRolesService authServicePort.IGetAllRoles
		ActionResult       actionResultServicePort.IActionResult
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
