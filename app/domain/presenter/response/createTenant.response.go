package responsePresenter

import "app/domain/model"

type (
	CreateTenantOutput struct {
		Tenant *model.Tenant `json:"tenant"`
		User   *model.User   `json:"user"`
	}

	CreateTenantResponse struct {
		Message string              `json:"message"`
		Data    *CreateTenantOutput `json:"data"`
		//Data    *model.Tenant `json:"data"`
	}
)
