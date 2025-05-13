package responsePresenter

import "app/model"

type (
	GetAllRolesResponse[Data_T any] struct {
		Message string       `json:"message"`
		Data    []model.Role `json:"data"`
	}
)
