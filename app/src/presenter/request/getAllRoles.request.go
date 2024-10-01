package requestPresenter

import (
	accessTokenServicePort "app/src/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	GetAllRolesRequest struct {
		ctx  iris.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetAllRolesRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *GetAllRolesRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *GetAllRolesRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *GetAllRolesRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
