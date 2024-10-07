package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GrantUserAsTenantAgent struct {
		TenantUUID *uuid.UUID `param:"tenantUUID" validate:"required"`
		UserUUID   *uuid.UUID `param:"userUUID" validate:"required"`
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GrantUserAsTenantAgent) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GrantUserAsTenantAgent) GetContext() context.Context {
	return this.ctx
}

func (this *GrantUserAsTenantAgent) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GrantUserAsTenantAgent) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
