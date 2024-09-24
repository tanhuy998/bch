package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"
	"context"
	"encoding/json"

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

	//_, err := this.CreateUserService.Serve(input.Data.Username, input.Data.Password, input.Data.Name, input.GetContext())

	newUser := &model.User{
		Username:   input.Data.Username,
		PassWord:   input.Data.Password,
		Name:       input.Data.Name,
		TenantUUID: *input.GetAuthority().GetTenantUUID(),
	}

	_, err := this.CreateUserService.CreateByModel(newUser, input.GetContext())

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	ret, err := this.GetSingleUserService.SearchByUsername(input.Data.Username, context.TODO())

	if err != nil {

		return nil, err
	}

	output.Message = "success"
	output.Data = ret

	resContent, err := json.Marshal(output)

	if err != nil {

		return nil, err
	}

	return this.ActionResult.Prepare().SetCode(201).SetContent(resContent), nil
}
