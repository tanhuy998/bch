package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	CreateUserPresenter struct {
		responseOutput.CreatedResponse
		Message string      `json:"message"`
		Data    *model.User `json:"data"`
	}
)
