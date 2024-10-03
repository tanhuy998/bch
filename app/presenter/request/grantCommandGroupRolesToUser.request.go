package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type (
	GrantCommandGroupRolesToUserRequest struct {
		GroupUUID *uuid.UUID  `param:"groupUUID"`
		UserUUID  *uuid.UUID  `param:"userUUID"`
		Data      []uuid.UUID `json:"data"`
		ctx       iris.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GrantCommandGroupRolesToUserRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *GrantCommandGroupRolesToUserRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *GrantCommandGroupRolesToUserRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GrantCommandGroupRolesToUserRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
