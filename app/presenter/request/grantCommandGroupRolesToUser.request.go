package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GrantCommandGroupRolesToUserRequest struct {
		GroupUUID *uuid.UUID  `param:"groupUUID"`
		UserUUID  *uuid.UUID  `param:"userUUID"`
		Data      []uuid.UUID `json:"data"`
		ctx       context.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GrantCommandGroupRolesToUserRequest) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *GrantCommandGroupRolesToUserRequest) GetContext() context.Context {

	return this.ctx
}

func (this *GrantCommandGroupRolesToUserRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GrantCommandGroupRolesToUserRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
