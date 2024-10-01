package responsePresenter

import (
	"app/src/model"
)

type (
	GetSingleAssignmentResponse struct {
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
