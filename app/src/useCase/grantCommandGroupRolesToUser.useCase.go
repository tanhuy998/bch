package usecase

import (
	actionResultService "app/service/actionResult"
	authService "app/service/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGrantCommandGroupRolesToUser interface {
		Execute(
			input *requestPresenter.GrantCommandGroupRolesToUserRequest,
			output *responsePresenter.GrantCommandGroupRolesToUserResponse,
		) (mvc.Result, error)
	}

	GrantCommandGroupRolesToUserUseCase struct {
		GrantCommandGroupRolesToUserService authService.IGrantCommandGroupRolesToUser
		ActionResult                        actionResultService.IActionResult
	}
)

func (this *GrantCommandGroupRolesToUserUseCase) Execute(
	input *requestPresenter.GrantCommandGroupRolesToUserRequest,
	output *responsePresenter.GrantCommandGroupRolesToUserResponse,
) (mvc.Result, error) {

	err := this.GrantCommandGroupRolesToUserService.Serve(input.GroupUUID, input.UserUUID, input.Data)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"

	rawContent, err := json.Marshal(output)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil
}
