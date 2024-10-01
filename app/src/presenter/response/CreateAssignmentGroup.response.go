package responsePresenter

import "app/src/model"

type (
	CreateAssignmentGroupResponse struct {
		Message string                 `json:"message"`
		Data    *model.AssignmentGroup `json:"data,omitempty"`
	}
)
