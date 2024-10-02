package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateAssignmentGroupResponse struct {
		responseOutput.CreatedResponse
		Message string                 `json:"message"`
		Data    *model.AssignmentGroup `json:"data,omitempty"`
	}
)
