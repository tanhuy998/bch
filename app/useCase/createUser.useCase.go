package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateUser interface {
		Execute(
			input *requestPresenter.CreateUserRequestPresenter,
			output *responsePresenter.CreateUserPresenter,
		) (mvc.Result, error)
	}

	CreateUserUsecase struct {
		CreateUserService    authService.ICreateUser
		ActionResult         actionResultService.IActionResult
		GetSingleUserService authService.IGetSingleUser
	}
)

func (this *CreateUserUsecase) Execute(
	input *requestPresenter.CreateUserRequestPresenter,
	output *responsePresenter.CreateUserPresenter,
) (mvc.Result, error) {

	_, err := this.CreateUserService.Serve(input.Data.Username, input.Data.Password, input.Data.Name)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	ret, err := this.GetSingleUserService.SearchByUsername(input.Data.Username)

	// if err != nil {

	// 	return nil, err
	// }

	output.Message = "success"
	output.Data = ret

	return this.ActionResult.ServeResponse(output)
}
