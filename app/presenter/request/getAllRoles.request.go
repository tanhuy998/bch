package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetAllRolesRequest struct {
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		tenantUUID uuid.UUID
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

func (this *GetAllRolesRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetAllRolesRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetAllRolesRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
