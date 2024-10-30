package responsePresenter

import (
	"app/model"
)

type (
	GetTenantUsers struct {
		Data []model.User `json:"data"`
	}
)
