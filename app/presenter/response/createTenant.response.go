package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateTenantOutput struct {
		Tenant *model.Tenant `json:"tenant"`
		User   *model.User   `json:"user"`
	}

	CreateTenantResponse struct {
		responseOutput.CreatedResponse
		Message string              `json:"message"`
		Data    *CreateTenantOutput `json:"data"`
		//Data    *model.Tenant `json:"data"`
	}
)
