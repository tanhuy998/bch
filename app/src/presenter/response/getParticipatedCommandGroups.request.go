package responsePresenter

import "app/src/valueObject"

type (
	GetParticipatedGroups struct {
		Message string                                      `json:"message"`
		Data    *valueObject.ParticipatedCommandGroupReport `json:"data"`
	}
)
