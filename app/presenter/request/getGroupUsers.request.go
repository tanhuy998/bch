package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetGroupUsersRequest struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		GroupUUID  *uuid.UUID `json:"groupUUID" validate:"required"`
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

func (this *GetGroupUsersRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GetGroupUsersRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GetGroupUsersRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
