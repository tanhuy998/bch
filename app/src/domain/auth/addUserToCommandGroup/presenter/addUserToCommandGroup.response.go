package responsePresenter

import "app/domain/model"

type (
	AddUserToCommandGroupResponse struct {
		Message string                  `json:"message"`
		Data    *model.CommandGroupUser `json:"data"`
	}
)
