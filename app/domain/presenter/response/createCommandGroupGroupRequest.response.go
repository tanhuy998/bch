package responsePresenter

import "github.com/google/uuid"

type (
	CreateCommandGroupResData struct {
		UUID uuid.UUID `json:"uuid"`
	}

	CreateCommandGroupResponse struct {
		Message string                    `json:"message"`
		Data    CreateCommandGroupResData `json:"data"`
	}
)
