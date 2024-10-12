package responsePresenter

import "app/valueObject"

type (
	ReportParticipatedGroups struct {
		Message string                                      `json:"message"`
		Data    *valueObject.ParticipatedCommandGroupReport `json:"data"`
	}
)
