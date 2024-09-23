package responsePresenter

import (
	"app/domain/model"
)

type (
	GetSingleAssignmentResponse struct {
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
