package getAllRoleDomain

import (
	actionResultServicePort "app/port/actionResult"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
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
		usecasePort.UseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse]
		GetAllRolesService authServicePort.IGetAllRoles
		ActionResult       actionResultServicePort.IActionResult
	}
)

func (this *GetAllRolesUseCase) Execute(
	input *requestPresenter.GetAllRolesRequest,
) (*responsePresenter.GetAllRolesResponse, error) {

	ret, err := this.GetAllRolesService.Serve(input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = ret

	return output, nil
}
