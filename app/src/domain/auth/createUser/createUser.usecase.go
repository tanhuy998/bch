package createUserDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
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
