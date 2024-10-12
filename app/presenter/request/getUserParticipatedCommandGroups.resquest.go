package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetUserParticipatedCommandGroups struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		UserUUID   *uuid.UUID `param:"userUUID" validate:"required"`
	}
)

func (this *GetUserParticipatedCommandGroups) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetUserParticipatedCommandGroups) GetContext() context.Context {

	return this.ctx
}

func (this *GetUserParticipatedCommandGroups) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetUserParticipatedCommandGroups) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *GetUserParticipatedCommandGroups) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetUserParticipatedCommandGroups) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetUserParticipatedCommandGroups) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
