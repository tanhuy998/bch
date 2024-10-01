package createTenantAgentDomain

import (
	actionResultServicePort "app/src/port/actionResult"
	tenantServicePort "app/src/port/tenant"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

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
		CreateTenantAgentService tenantServicePort.ICreateTenantAgent
		ActionResult             actionResultServicePort.IActionResult
	}
)

// func (this *CreateTenantAgentUseCase) Execute(
// 	input *requestPresenter.CreateTenantAgentRequest,
// 	output *responsePresenter.CreateTenantAgentResponse,
// ) (mvc.Result, error) {

// 	newAgent, err := this.CreateTenantAgentService.Serve(input.Data.Username, input.Data.Password, input.Data.Name, input.GetContext())

// 	if err != nil {

// 		return this.ActionResult.ServeErrorResponse(err)
// 	}

// 	output.Message = "success"
// 	output.Data = newAgent
// 	rawContent, err := json.Marshal(output)

// 	if err != nil {

// 		return this.ActionResult.ServeErrorResponse(err)
// 	}

// 	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent).Done()
// }
