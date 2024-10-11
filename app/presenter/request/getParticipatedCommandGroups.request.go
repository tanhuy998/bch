package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetParticipatedGroups struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		UserUUID   *uuid.UUID `param:"userUUID" validate:"required"`
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

func (this *GetParticipatedGroups) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetParticipatedGroups) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetParticipatedGroups) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
