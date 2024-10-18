package responsePresenter

import "app/model"

type (
	GetAssignmentGroupUnAssignedCommandGroupUsers struct {
		Message string                    `json:"message"`
		Data    []*model.CommandGroupUser `json:"data"`
	}
)
