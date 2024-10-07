package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"
)

type (
	GetAllRolesRequest struct {
		ctx  context.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetAllRolesRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetAllRolesRequest) GetContext() context.Context {

	return this.ctx
}

func (this *GetAllRolesRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetAllRolesRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
