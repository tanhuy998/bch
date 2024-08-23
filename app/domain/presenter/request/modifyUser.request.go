package requestPresenter

type (
	InputModifyUser struct {
		Name     string `json:"name" validate:"required_without_all=Password"`
		Password string `json:"password" validate:"required_without_all=Name"`
	}

	ModifyUserRequest struct {
		UserUUID string           `param:"userUUID" validate:"required"`
		Data     *InputModifyUser `json:"data" validate:"required"`
	}
)
