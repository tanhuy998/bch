package responsePresenter

import (
	"app/internal/responseOutput"
	"app/model"
)

type (
	AddUserToCommandGroupResponse struct {
		responseOutput.CreatedResponse
		Message string                  `json:"message"`
		Data    *model.CommandGroupUser `json:"data"`
	}
)
