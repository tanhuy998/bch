package responsePresenter

import "app/model"

type (
	GetAllRolesResponse struct {
		Message string        `json:"message"`
		Data    []*model.Role `json:"data"`
	}
)
