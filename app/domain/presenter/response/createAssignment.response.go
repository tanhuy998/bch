package responsePresenter

import (
	"app/domain/model"
)

type (
	CreateAssignmentResponse struct {
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
