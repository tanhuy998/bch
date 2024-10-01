package requestPresenter

import (
	accessTokenServicePort "app/src/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	GrantCommandGroupRolesToUserRequest struct {
		GroupUUID string   `param:"groupUUID"`
		UserUUID  string   `param:"userUUID"`
		Data      []string `json:"data"`
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
