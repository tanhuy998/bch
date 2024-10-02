package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateAssignmentResponse struct {
		responseOutput.CreatedResponse
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
