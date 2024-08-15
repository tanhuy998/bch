package requestPresenter

import "app/domain/model"

type (
	CreateCommandGroupRequest struct {
		Data model.CommandGroup `json:"data" validate:"required"`
	}
)
