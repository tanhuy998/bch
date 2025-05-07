package responsePresenter

type (
	GetUserParticipatedCommandGroups[Entity_T any] struct {
		Message string     `json:"message"`
		Data    []Entity_T `json:"data"`
	}
)
