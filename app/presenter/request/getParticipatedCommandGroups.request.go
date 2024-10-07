package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetParticipatedGroups struct {
		UserUUID *uuid.UUID `param:"userUUID" validate:"required"`
		ctx      context.Context
		auth     accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetParticipatedGroups) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetParticipatedGroups) GetContext() context.Context {

	return this.ctx
}

func (this *GetParticipatedGroups) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetParticipatedGroups) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
