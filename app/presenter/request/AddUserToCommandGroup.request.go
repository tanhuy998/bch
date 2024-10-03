package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type (
	data_AddUserToCommandGroup struct {
		RoleUUIDs []string `json:"roleUUIDs"`
	}

	AddUserToCommandGroupRequest struct {
		GroupUUID *uuid.UUID                 `param:"groupUUID" validate:"required"`
		UserUUID  *uuid.UUID                 `param:"userUUID" validate:"required"`
		Data      data_AddUserToCommandGroup `json:"data"`
		ctx       iris.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *AddUserToCommandGroupRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *AddUserToCommandGroupRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *AddUserToCommandGroupRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *AddUserToCommandGroupRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
