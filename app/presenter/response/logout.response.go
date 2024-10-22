package responsePresenter

import (
	"app/internal/responseOutput"
)

type (
	Logout struct {
		responseOutput.HTTPStatusResponse
		Message string `json:"message"`
	}
)
