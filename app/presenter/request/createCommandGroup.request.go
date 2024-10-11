package requestPresenter

import (
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateCommandGroupRequest struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		Data       *model.CommandGroup `json:"data" validate:"required"`
	}
)

func (this *CreateCommandGroupRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateCommandGroupRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateCommandGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateCommandGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *CreateCommandGroupRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateCommandGroupRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *CreateCommandGroupRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
