package responsePresenter

import "app/model"

type (
	AuthNavigateTenant struct {
		Message string          `json:"message"`
		Data    []*model.Tenant `json:"data"`
	}
)
