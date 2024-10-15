package responsePresenter

import "app/model"

type (
	GetAssignments struct {
		Data    []model.Assignment `json:"data"`
		Message string             `json:"message"`
	}
)
