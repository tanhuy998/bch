package responsePresenter

import "app/model"

type (
	GetTenantAllGroups struct {
		Message string                `json:"message"`
		Data    []*model.CommandGroup `json:"data"`
	}
)
