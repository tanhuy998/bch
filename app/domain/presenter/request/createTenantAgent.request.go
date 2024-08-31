package requestPresenter

import "app/domain/model"

type (
	CreateTenantAgentRequest struct {
		Data *model.TenantAgent `json:"data" validate:"required"`
	}
)
