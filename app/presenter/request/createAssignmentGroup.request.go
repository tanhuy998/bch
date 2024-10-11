package requestPresenter

import (
	"app/model"
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateAssignmentGroupRequest struct {
		tenantUUID     uuid.UUID
		AssignmentUUID *uuid.UUID             `param:"assignmnetUUID" validate:"required"`
		Data           *model.AssignmentGroup `json:"data" validate:"required"`
		ctx            context.Context
		auth           accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateAssignmentGroupRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateAssignmentGroupRequest) GetContext() context.Context {

	return this.ctx
}

func (this *CreateAssignmentGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateAssignmentGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *CreateAssignmentGroupRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateAssignmentGroupRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *CreateAssignmentGroupRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
