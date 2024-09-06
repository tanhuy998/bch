package responsePresenter

import "app/domain/model"

type (
	CreateTenantResponse struct {
		Message string        `json:"message"`
		Data    *model.Tenant `json:"data"`
	}
)
