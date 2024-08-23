package requestPresenter

type (
	InputUser struct {
		//UUID          uuid.UUID `json:"uuid" bson:"uuid"`
		Name     string `json:"name" bson:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		//IsDeactivated bool   `json:"deactivated" bson:"deactivated"`
		//Info          UserInfo  `json:"userInfo" bson:"userInfo"`
	}

	CreateUserRequestPresenter struct {
		Data *InputUser `json:"data" validate:"required"`
	}
)
