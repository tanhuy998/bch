package responsePresenter

import "app/src/model"

type (
	CreateUserPresenter struct {
		Message string      `json:"message"`
		Data    *model.User `json:"data"`
	}
)
