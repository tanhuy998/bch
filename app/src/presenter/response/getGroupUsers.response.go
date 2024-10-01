package responsePresenter

import "app/src/model"

type (
	GetGroupUsersResponse struct {
		Message string        `json:"message"`
		Data    []*model.User `json:"data"`
	}
)
