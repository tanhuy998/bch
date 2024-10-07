package createTenantAgentDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	"github.com/google/uuid"
)

type (
	// ICreateTenantAgent interface {
	// 	Execute(
	// 		input *requestPresenter.CreateTenantAgentRequest,
	// 		output *responsePresenter.CreateTenantAgentResponse,
	// 	) (mvc.Result, error)
	// }

	CreateTenantAgentUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateTenantAgentRequest, responsePresenter.CreateTenantAgentResponse]
		CreateTenantAgentService tenantServicePort.ICreateTenantAgent
	}
)

func (this *CreateTenantAgentUseCase) Execute(
	input *requestPresenter.CreateTenantAgentRequest,
) (*responsePresenter.CreateTenantAgentResponse, error) {

	newUser := &model.User{
		Name:     input.Data.Name,
		Username: input.Data.Username,
		PassWord: input.Data.Password,
	}

	auth := input.GetAuthority()

	if auth != nil {

		newUser.CreatedBy = libCommon.PointerPrimitive(auth.GetUserUUID())
	}

	newAgentModel := &model.TenantAgent{
		UUID:       libCommon.PointerPrimitive(uuid.New()),
		UserUUID:   libCommon.PointerPrimitive(uuid.UUID(*newUser.UUID)),
		TenantUUID: input.TenantUUID,
		//Deactivated: true,
	}

	_, tenantAgent, err := this.CreateTenantAgentService.Serve(newUser, newAgentModel, *input.TenantUUID, input.GetContext())

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = tenantAgent

	return output, nil
}
