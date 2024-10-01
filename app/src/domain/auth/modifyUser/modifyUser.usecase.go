package modifyUserDomain

import (
	"app/src/internal/common"
	"app/src/model"
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"encoding/json"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

type (
	// IModifyUser interface {
	// 	Execute(
	// 		input *requestPresenter.ModifyUserRequest,
	// 		output *responsePresenter.ModifyUserResponse,
	// 	) (mvc.Result, error)
	// }

	ModifyUserUseCase struct {
		ModifyUser   authServicePort.IModifyUser
		ActionResult actionResultServicePort.IActionResult
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

	if errors.Is(err, common.ERR_NOT_FOUND) {

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
