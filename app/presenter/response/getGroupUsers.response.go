package responsePresenter

import (
	"app/model"
)

type (
	GetGroupUsersResponse struct {
		Message string                    `json:"message"`
		Data    []*model.CommandGroupUser `json:"data"`
	}
)
