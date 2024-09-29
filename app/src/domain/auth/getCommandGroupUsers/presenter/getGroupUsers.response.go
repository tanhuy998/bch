package responsePresenter

import "app/domain/model"

type (
	GetGroupUsersResponse struct {
		Message string        `json:"message"`
		Data    []*model.User `json:"data"`
	}
)
