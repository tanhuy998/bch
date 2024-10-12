package responsePresenter

import (
	"app/model"
)

type (
	GetUserParticipatedCommandGroups struct {
		Message string                `json:"message"`
		Data    []*model.CommandGroup `json:"data"`
	}
)
