package requestPresenter

import (
	"context"

	"github.com/google/uuid"
)

type (
	SwitchTenant struct {
		TenantUUID *uuid.UUID
		ctx        context.Context
	}
)

func (this *SwitchTenant) ReceiveContext(ctx context.Context) {

	this.ctx = ctx
}

func (this *SwitchTenant) GetContext() context.Context {

	return this.ctx
}
