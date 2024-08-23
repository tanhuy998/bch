package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetGroupUsers interface {
		Execute(
			input *requestPresenter.GetGroupUsersRequest,
			output *responsePresenter.GetGroupUsersResponse,
		) (mvc.Result, error)
	}

	GetGroupUsersUseCase struct {
		GetCommandGroupUsersService authService.IGetCommandGroupUsers
		ActionResult                actionResultService.IActionResult
	}
)

func (this *GetGroupUsersUseCase) Execute(
	input *requestPresenter.GetGroupUsersRequest,
	output *responsePresenter.GetGroupUsersResponse,
) (mvc.Result, error) {

	data, err := this.GetCommandGroupUsersService.Serve(input.GroupUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data = data

	return this.ActionResult.ServeResponse(output)
}
