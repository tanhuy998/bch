package requestPresenter

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	CreateTenantInputData struct {
		Name        string     `json:"name" validate:"required"`
		Description string     `json:"description"`
		User        *InputUser `json:"user"`
		//TenantAgentUUID string `json:"tenantAgentUUID"`
	}

	CreateTenantRequest struct {
		Data CreateTenantInputData `json:"data"`
		ctx  iris.Context
		auth accessTokenServicePort.IAccessTokenAuthData
	}
)

func (this *CreateTenantRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *CreateTenantRequest) GetContext() iris.Context {

	return this.ctx
}

func (this *CreateTenantRequest) GetAuthority() accessTokenServicePort.IAccessTokenAuthData {

	return this.auth
}

func (this *CreateTenantRequest) SetAuthority(auth accessTokenServicePort.IAccessTokenAuthData) {

	this.auth = auth
}
