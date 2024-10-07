package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"
)

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
		ctx  context.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateUserRequestPresenter) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateUserRequestPresenter) GetContext() context.Context {

	return this.ctx
}

func (this *CreateUserRequestPresenter) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateUserRequestPresenter) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
