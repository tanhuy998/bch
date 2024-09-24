package requestPresenter

import (
	accessTokenServicePort "app/adapter/accessToken"

	"github.com/kataras/iris/v12"
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
		ctx  iris.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateUserRequestPresenter) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *CreateUserRequestPresenter) GetContext() iris.Context {

	return this.ctx
}

func (this *CreateUserRequestPresenter) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateUserRequestPresenter) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
