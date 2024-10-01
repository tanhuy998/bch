package requestPresenter

import (
	"app/src/model"
	accessTokenServicePort "app/src/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	CreateCommandGroupRequest struct {
		Data *model.CommandGroup `json:"data" validate:"required"`
		ctx  iris.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateCommandGroupRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *CreateCommandGroupRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *CreateCommandGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateCommandGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
