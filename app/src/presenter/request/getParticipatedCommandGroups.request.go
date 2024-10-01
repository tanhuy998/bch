package requestPresenter

import (
	accessTokenServicePort "app/src/port/accessToken"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type (
	GetParticipatedGroups struct {
		UserUUID *uuid.UUID `param:"userUUID" validate:"required"`
		ctx      iris.Context
		auth     accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetParticipatedGroups) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *GetParticipatedGroups) GetContext() iris.Context {

	return this.ctx
}

func (this *GetParticipatedGroups) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetParticipatedGroups) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
