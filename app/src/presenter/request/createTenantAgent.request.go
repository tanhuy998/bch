package requestPresenter

import (
	"github.com/kataras/iris/v12"
)

type (
	CreateTenantAgentRequest struct {
		//Data *model.TenantAgent `json:"data" validate:"required"`
		Data *InputUser `json:"data" validate:"required"`
		ctx  iris.Context
	}
)

func (this *CreateTenantAgentRequest) ReceiveContext(ctx iris.Context) {

	this.ctx = ctx
}

func (this *CreateTenantAgentRequest) GetContext() iris.Context {

	return this.ctx
}
