package requestPresenter

import (
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	"context"
)

type (
	CreateCommandGroupRequest struct {
		Data *model.CommandGroup `json:"data" validate:"required"`
		ctx  context.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateCommandGroupRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateCommandGroupRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateCommandGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateCommandGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
