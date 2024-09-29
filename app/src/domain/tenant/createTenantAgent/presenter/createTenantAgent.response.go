package responsePresenter

import "app/domain/model"

type (
	CreateTenantAgentResponse struct {
		Message string             `json:"message"`
		Data    *model.TenantAgent `json:"data,omitempty"`
	}
)
