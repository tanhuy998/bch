package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"
	"context"

	"github.com/google/uuid"
)

type (
	GrantCommandGroupRolesToUserRequest struct {
		tenantUUID uuid.UUID
		GroupUUID  *uuid.UUID  `param:"groupUUID"`
		UserUUID   *uuid.UUID  `param:"userUUID"`
		Data       []uuid.UUID `json:"data"`
		ctx        context.Context
		auth       accessTokenServicePort.IAccessTokenAuthData
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

func (this *GrantCommandGroupRolesToUserRequest) SetTenantUUID(tenantUUID uuid.UUID) {

	this.tenantUUID = tenantUUID
}

func (this *GrantCommandGroupRolesToUserRequest) IsValidTenantUUID() bool {

	return this.tenantUUID != uuid.Nil
}

func (this *GrantCommandGroupRolesToUserRequest) GetTenantUUID() uuid.UUID {

	return this.tenantUUID
}
