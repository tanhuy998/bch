package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateTenantAgentResponse struct {
		responseOutput.CreatedResponse
		Message string             `json:"message"`
		Data    *model.TenantAgent `json:"data,omitempty"`
	}
)
