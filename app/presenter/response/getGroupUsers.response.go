package responsePresenter

import (
	"app/model"
)

type (
	GetGroupUsersResponse struct {
		Message string        `json:"message"`
		Data    []*model.User `json:"data"`
	}
)
