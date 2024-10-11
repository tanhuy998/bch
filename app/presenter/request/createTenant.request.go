package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateTenantInputData struct {
		Title       string     `json:"title" validate:"required"`
		Description string     `json:"description"`
		User        *InputUser `json:"user"`
		//TenantAgentUUID string `json:"tenantAgentUUID"`
	}

	CreateTenantRequest struct {
		Data       CreateTenantInputData `json:"data"`
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateTenantRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateTenantRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateTenantRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateTenantRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *CreateTenantRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateTenantRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}

func (this *CreateTenantRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}
