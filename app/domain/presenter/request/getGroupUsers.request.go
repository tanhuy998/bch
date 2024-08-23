package requestPresenter

type (
	GetGroupUsersRequest struct {
		GroupUUID string `json:"groupUUID" validate:"required"`
	}
)
