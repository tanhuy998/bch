package getAllRoleDomain

import (
	"app/model"
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
			output *responsePresenter.GetAllRolesResponse[model.Role],
		) (mvc.Result, error)
	}

	GetAllRolesUseCase struct {
		usecasePort.UseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse[model.Role]]
		GetAllRolesService authServicePort.IGetAllRoles
		//ActionResult       actionResultServicePort.IActionResult
	}
)

func (this *GetAllRolesUseCase) Execute(
	input *requestPresenter.GetAllRolesRequest,
) (*responsePresenter.GetAllRolesResponse[model.Role], error) {

	ret, err := this.GetAllRolesService.Serve(input.GetContext())

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = ret

	return output, nil
}
