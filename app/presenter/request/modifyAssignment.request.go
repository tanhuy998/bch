package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	modifyAssignmentInput struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Deadline    *time.Time `json:"deadline"`
	}
	ModifyAssignment struct {
		tenantUUID     uuid.UUID
		ctx            context.Context
		auth           accessTokenServicePort.IAccessTokenAuthData
		AssignmentUUID *uuid.UUID             `param:"assignmentUUID" validate:"required"`
		Data           *modifyAssignmentInput `json:"data" validate:"required"`
	}
)

func (this *ModifyAssignment) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *ModifyAssignment) GetContext() context.Context {

	return this.ctx
}

func (this *ModifyAssignment) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *ModifyAssignment) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *ModifyAssignment) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *ModifyAssignment) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *ModifyAssignment) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
