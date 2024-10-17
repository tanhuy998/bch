package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	CreateAssignmentGroupMember struct {
		tenantUUID              uuid.UUID
		ctx                     context.Context
		auth                    accessTokenServicePort.IAccessTokenAuthData
		AssignmentGroupUUID     *uuid.UUID  `param:"groupUUID" validate:"required"`
		ComandGroupUserUUIDList []uuid.UUID `json:"data" validate:"required"`
	}
)

func (this *CreateAssignmentGroupMember) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *CreateAssignmentGroupMember) GetContext() context.Context {

	return this.ctx
}

func (this *CreateAssignmentGroupMember) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateAssignmentGroupMember) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *CreateAssignmentGroupMember) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *CreateAssignmentGroupMember) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *CreateAssignmentGroupMember) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
