package responsePresenter

import "app/valueObject"

type (
	GetParticipatedGroups struct {
		Message string                                      `json:"message"`
		Data    *valueObject.ParticipatedCommandGroupReport `json:"data"`
	}
)
