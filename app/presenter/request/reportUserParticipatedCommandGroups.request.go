package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	ReportParticipatedGroups struct {
		tenantUUID uuid.UUID
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
		UserUUID   *uuid.UUID `param:"userUUID" validate:"required"`
	}
)

func (this *ReportParticipatedGroups) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *ReportParticipatedGroups) GetContext() context.Context {

	return this.ctx
}

func (this *ReportParticipatedGroups) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *ReportParticipatedGroups) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}

func (this *ReportParticipatedGroups) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *ReportParticipatedGroups) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *ReportParticipatedGroups) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
