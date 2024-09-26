package requestPresenter

import (
	accessTokenServicePort "app/adapter/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	GetGroupUsersRequest struct {
		GroupUUID string `json:"groupUUID" validate:"required"`
		ctx       iris.Context
		auth      accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetGroupUsersRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *GetGroupUsersRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *GetGroupUsersRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetGroupUsersRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
