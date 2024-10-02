package createUserDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	actionResultServicePort "app/port/actionResult"
	authServicePort "app/port/auth"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
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
		CreateUserService    authServicePort.ICreateUser
		ActionResult         actionResultServicePort.IActionResult
		GetSingleUserService authServicePort.IGetSingleUser
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
		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
		CreatedBy:  libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
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
