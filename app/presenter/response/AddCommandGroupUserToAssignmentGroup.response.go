package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateAssignmentGroupMemeber struct {
		responseOutput.AccepptedReponse
		Message string `json:"message"`
		Data    []*model.AssignmentGroupMember
	}
)
