package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetGroupUsersRequest struct {
		GroupUUID *uuid.UUID `json:"groupUUID" validate:"required"`
		ctx       context.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetGroupUsersRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetGroupUsersRequest) GetContext() context.Context {

	return this.ctx
}

func (this *GetGroupUsersRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetGroupUsersRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
