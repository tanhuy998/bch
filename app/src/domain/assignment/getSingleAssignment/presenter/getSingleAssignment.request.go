package requestPresenter

import (
	accessTokenServicePort "app/adapter/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	GetSingleAssignmentRequest struct {
		UUID      string `param:"uuid"`
		ctx       iris.Context
		authority accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *GetSingleAssignmentRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.authority
}

func (this *GetSingleAssignmentRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.authority = auth
}

func (this *GetSingleAssignmentRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *GetSingleAssignmentRequest) GetContext() iris.Context {

	return this.ctx
}
