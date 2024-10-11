package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateAssignmentInput struct {
		TenantUUID *uuid.UUID `json:"tenantUUID" bson:"tenantUUID"`
		Title      string     `json:"title" bson:"title" validate:"required"`
	}

	CreateAssigmentRequest struct {
		tenantUUID uuid.UUID
		authority  accessTokenServicePort.IAccessTokenAuthData
		ctx        context.Context
		Data       *CreateAssignmentInput `json:"data,omitempty" validate:"required"`
	}
)

func (this *CreateAssigmentRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.authority
}

func (this *CreateAssigmentRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.authority = auth
}

func (this *CreateAssigmentRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateAssigmentRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateAssigmentRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateAssigmentRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *CreateAssigmentRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
