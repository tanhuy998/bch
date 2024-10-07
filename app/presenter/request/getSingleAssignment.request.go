package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GetSingleAssignmentRequest struct {
		UUID      *uuid.UUID `param:"uuid"`
		ctx       context.Context
		authority accessTokenServicePort.IAccessTokenAuthData
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
