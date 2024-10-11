package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetSingleAssignmentRequest struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		authority  accessTokenServicePort.IAccessTokenAuthData
		UUID       *uuid.UUID `param:"uuid"`
	}
)

func (this *GetSingleAssignmentRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.authority
}

func (this *GetSingleAssignmentRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.authority = auth
}

func (this *GetSingleAssignmentRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GetSingleAssignmentRequest) GetContext() context.Context {

	return this.ctx
}

func (this *GetSingleAssignmentRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetSingleAssignmentRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetSingleAssignmentRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
