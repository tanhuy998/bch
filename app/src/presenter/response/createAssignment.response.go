package responsePresenter

import (
	"app/src/model"
)

type (
	CreateAssignmentResponse struct {
		Message string            `json:"message"`
		Data    *model.Assignment `json:"data,omitempty"`
	}
)
