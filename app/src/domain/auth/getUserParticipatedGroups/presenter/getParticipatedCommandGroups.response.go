package responsePresenter

import "app/domain/valueObject"

type (
	GetParticipatedGroups struct {
		Message string                                      `json:"message"`
		Data    *valueObject.ParticipatedCommandGroupReport `json:"data"`
	}
)
