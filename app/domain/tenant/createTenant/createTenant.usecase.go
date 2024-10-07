package createTenantDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	// ICreateTenant interface {
	// 	Execute(
	// 		input *requestPresenter.CreateTenantRequest,
	// 		output *responsePresenter.CreateTenantResponse,
	// 	) (mvc.Result, error)
	// }

	CreateTenantUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse]
		CreateTenantService tenantServicePort.ICreateTenant
		//ActionResult        actionResultServicePort.IActionResult
	}
)

func (this *CreateTenantUseCase) Execute(
	input *requestPresenter.CreateTenantRequest,
) (*responsePresenter.CreateTenantResponse, error) {

	newTenant := &model.Tenant{
		Name:        input.Data.Name,
		Description: input.Data.Description,
	}

	auth := input.GetAuthority()

	if auth != nil {

		newTenant.CreatedBy = libCommon.PointerPrimitive(auth.GetUserUUID())
	}

	var newUser *model.User

	if input.Data.User != nil {

		newUser = &model.User{
			Name:     input.Data.User.Name,
			Username: input.Data.User.Username,
			PassWord: input.Data.User.Password,
		}
	}

	newTenant, newUser, err := this.CreateTenantService.Serve(newTenant, newUser, input.GetContext())

	if err != nil {

		//return this.ActionResult.ServeErrorResponse(err)

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = &responsePresenter.CreateTenantOutput{
		Tenant: newTenant,
		User:   newUser,
	}

	// rawContent, _ := json.Marshal(output)

	// return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil

	return output, nil
}
