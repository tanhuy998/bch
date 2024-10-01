package usecase

import (
	"app/domain/model"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IModifyUser interface {
		Execute(
			input *requestPresenter.ModifyUserRequest,
			output *responsePresenter.ModifyUserResponse,
		) (mvc.Result, error)
	}

	ModifyUserUseCase struct {
		ModifyUser   authService.IModifyUser
		ActionResult actionResultService.IActionResult
	}
)

func (this *ModifyUserUseCase) Execute(
	input *requestPresenter.ModifyUserRequest,
	output *responsePresenter.ModifyUserResponse,
) (mvc.Result, error) {

	dataModel := &model.User{
		Name:     input.Data.Name,
		PassWord: input.Data.Password,
	}

	err := this.ModifyUser.Serve(input.UserUUID, dataModel)

	if err == authService.ERR_MODIFY_USER_NOT_FOUND {

		output.Message = err.Error()

		rawContent, _ := json.Marshal(output)

		return this.ActionResult.Prepare().SetCode(404).SetContent(rawContent), nil
	}

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResult.ServeResponse(output)
}
