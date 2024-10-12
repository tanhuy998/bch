package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetTenantAllGroups struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetTenantAllGroups) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetTenantAllGroups) GetContext() context.Context {

	return this.ctx
}

func (this *GetTenantAllGroups) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetTenantAllGroups) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *GetTenantAllGroups) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetTenantAllGroups) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetTenantAllGroups) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
