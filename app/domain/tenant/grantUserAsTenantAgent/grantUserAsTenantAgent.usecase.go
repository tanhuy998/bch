package grantUserAsTenantAgentDomain

import (
	libCommon "app/internal/lib/common"
	"app/model"
	tenantServicePort "app/port/tenant"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	GrantUserAsTenantAgentUseCase struct {
		usecasePort.UseCase[requestPresenter.GrantUserAsTenantAgent, responsePresenter.GrantUserAsTenantAgent]
		GrantUserAsTenantAgentService tenantServicePort.IGrantUserAsTenantAgent
	}
)

func (this *GrantUserAsTenantAgentUseCase) Execute(
	input *requestPresenter.GrantUserAsTenantAgent,
) (*responsePresenter.GrantUserAsTenantAgent, error) {

	tenantAgentModel := &model.TenantAgent{
		TenantUUID: input.TenantUUID,
		UserUUID:   input.UserUUID,
	}

	auth := input.GetAuthority()

	if auth != nil {

		tenantAgentModel.CreatedBy = libCommon.PointerPrimitive(auth.GetUserUUID())
	}

	_, err := this.GrantUserAsTenantAgentService.Serve(
		*input.UserUUID, *input.TenantUUID, tenantAgentModel, input.GetContext(),
	)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()

	output.Message = "success"

	return output, nil
}
