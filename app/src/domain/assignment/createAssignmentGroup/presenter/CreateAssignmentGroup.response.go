package responsePresenter

import "app/domain/model"

type (
	CreateAssignmentGroupResponse struct {
		Message string                 `json:"message"`
		Data    *model.AssignmentGroup `json:"data,omitempty"`
	}
)
