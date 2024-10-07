package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	data_AddUserToCommandGroup struct {
		RoleUUIDs []string `json:"roleUUIDs"`
	}

	AddUserToCommandGroupRequest struct {
		GroupUUID *uuid.UUID                 `param:"groupUUID" validate:"required"`
		UserUUID  *uuid.UUID                 `param:"userUUID" validate:"required"`
		Data      data_AddUserToCommandGroup `json:"data"`
		ctx       context.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *AddUserToCommandGroupRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *AddUserToCommandGroupRequest) GetContext() context.Context {

	return this.ctx
}

func (this *AddUserToCommandGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *AddUserToCommandGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
