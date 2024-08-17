package requestPresenter

type (
	GetParticipatedGroups struct {
		UserUUID string `param:"userUUID" validate:"required"`
	}
)
