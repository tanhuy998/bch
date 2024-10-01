package usecase

import (
	tenantServicePort "app/adapter/tenant"
	"app/domain/model"
	actionResultService "app/service/actionResult"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateTenant interface {
		Execute(
			input *requestPresenter.CreateTenantRequest,
			output *responsePresenter.CreateTenantResponse,
		) (mvc.Result, error)
	}

	CreateTenantUseCase struct {
		CreateTenantService tenantServicePort.ICreateTenant
		ActionResult        actionResultService.IActionResult
	}
)

func (this *CreateTenantUseCase) Execute(
	input *requestPresenter.CreateTenantRequest,
	output *responsePresenter.CreateTenantResponse,
) (mvc.Result, error) {

	newTenant := &model.Tenant{
		Name:        input.Data.Name,
		Description: input.Data.Description,
	}

	newUser := &model.User{
		Name:     input.Data.User.Name,
		Username: input.Data.User.Username,
		PassWord: input.Data.User.Password,
	}

	newTenant, newUser, err := this.CreateTenantService.Serve(newTenant, newUser, input.GetContext())

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = &responsePresenter.CreateTenantOutput{
		Tenant: newTenant,
		User:   newUser,
	}

	rawContent, _ := json.Marshal(output)

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil
}
