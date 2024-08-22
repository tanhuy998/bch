package responsePresenter

import "app/domain/model"

type (
	GetAllRolesResponse struct {
		Message string        `json:"message"`
		Data    []*model.Role `json:"data"`
	}
)
