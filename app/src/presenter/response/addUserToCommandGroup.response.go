package responsePresenter

import "app/src/model"

type (
	AddUserToCommandGroupResponse struct {
		Message string                  `json:"message"`
		Data    *model.CommandGroupUser `json:"data"`
	}
)
