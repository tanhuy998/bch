package responsePresenter

import "app/src/model"

type (
	GetAllRolesResponse struct {
		Message string        `json:"message"`
		Data    []*model.Role `json:"data"`
	}
)
