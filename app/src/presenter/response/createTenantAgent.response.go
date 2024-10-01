package responsePresenter

import "app/src/model"

type (
	CreateTenantAgentResponse struct {
		Message string             `json:"message"`
		Data    *model.TenantAgent `json:"data,omitempty"`
	}
)
