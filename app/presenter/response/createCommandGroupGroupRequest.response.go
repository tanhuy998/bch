package responsePresenter

import (
	"app/internal/responseOutput"

	"github.com/google/uuid"
)

type (
	CreateCommandGroupResData struct {
		UUID uuid.UUID `json:"uuid"`
	}

	CreateCommandGroupResponse struct {
		responseOutput.CreatedResponse
		Message string                    `json:"message"`
		Data    CreateCommandGroupResData `json:"data"`
	}
)
