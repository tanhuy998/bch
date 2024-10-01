package getAllRoleDomain

import (
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

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
