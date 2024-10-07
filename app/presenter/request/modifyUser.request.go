package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	InputModifyUser struct {
		Name     string `json:"name" validate:"required_without_all=Password"`
		Password string `json:"password" validate:"required_without_all=Name"`
	}

	ModifyUserRequest struct {
		UserUUID *uuid.UUID       `param:"userUUID" validate:"required"`
		Data     *InputModifyUser `json:"data" validate:"required"`
		ctx      context.Context
		auth     accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *ModifyUserRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *ModifyUserRequest) GetContext() context.Context {

	return this.ctx
}

func (this *ModifyUserRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *ModifyUserRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
