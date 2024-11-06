package requestPresenter

import (
	"app/valueObject/requestInput"

	"github.com/google/uuid"
)

type (
	SwitchTenant struct {
		TenantUUID *uuid.UUID `param:"tenantUUID" validate:"required"`
		//ctx        context.Context
		requestInput.ContextInput
	}
)

// func (this *SwitchTenant) ReceiveContext(ctx context.Context) {

// 	this.ctx = ctx
// }

// func (this *SwitchTenant) GetContext() context.Context {

// 	return this.ctx
// }
