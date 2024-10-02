package responsePresenter

import (
	"app/model"
)

type (
	GetSingleAssignmentResponse struct {
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
