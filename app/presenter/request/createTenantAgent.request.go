package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateTenantAgentRequest struct {
		//Data *model.TenantAgent `json:"data" validate:"required"`
		TenantUUID *uuid.UUID `param:"tenantUUID" validate:"required"`
		Data       *InputUser `json:"data" validate:"required"`
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateTenantAgentRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateTenantAgentRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateTenantAgentRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateTenantAgentRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
