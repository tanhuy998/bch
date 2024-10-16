package responsePresenter

import "app/model"

type (
	GetAssignmentGroups struct {
		Message string                   `json:"message"`
		Data    []*model.AssignmentGroup `json:"data"`
	}
)
