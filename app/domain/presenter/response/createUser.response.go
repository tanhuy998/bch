package responsePresenter

import "app/domain/model"

type (
	CreateUserPresenter struct {
		Message string      `json:"message"`
		Data    *model.User `json:"data"`
	}
)
