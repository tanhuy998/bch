package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	tenantAgentService "app/service/tenantAgent"
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateTenantAgent interface {
		Execute(
			input *requestPresenter.CreateTenantAgentRequest,
			output *responsePresenter.CreateTenantAgentResponse,
		) (mvc.Result, error)
	}

	CreateTenantAgentUseCase struct {
		CreateTenantAgentService tenantAgentService.ICreaateTenantAgent
		ActionResult             actionResultService.IActionResult
	}
)

func (this *CreateTenantAgentUseCase) Execute(
	input *requestPresenter.CreateTenantAgentRequest,
	output *responsePresenter.CreateTenantAgentResponse,
) (mvc.Result, error) {

	err := this.CreateTenantAgentService.Serve(input.Data)

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
