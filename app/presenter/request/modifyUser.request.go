package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	InputModifyUser struct {
		Name     string `json:"name" validate:"required_without_all=Password"`
		Password string `json:"password" validate:"required_without_all=Name"`
	}

	ModifyUserRequest struct {
		UserUUID string           `param:"userUUID" validate:"required"`
		Data     *InputModifyUser `json:"data" validate:"required"`
		ctx      iris.Context
		auth     accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *ModifyUserRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *ModifyUserRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *ModifyUserRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *ModifyUserRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
