package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	tenantService "app/service/tenant"
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
		CreateTenantService tenantService.ICreateTenant
		ActionResult        actionResultService.IActionResult
	}
)

func (this *CreateTenantUseCase) Execute(
	input *requestPresenter.CreateTenantRequest,
	output *responsePresenter.CreateTenantResponse,
) (mvc.Result, error) {

	newTenant, err := this.CreateTenantService.Serve(input.Data.Name, input.Data.Description, input.Data.TenantAgentUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = newTenant

	rawContent, _ := json.Marshal(output)

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil
}
