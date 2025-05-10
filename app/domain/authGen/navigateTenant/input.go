package navigateTenantDomain

import (
	"context"

	"github.com/google/uuid"
)

type (
	domain_input struct {
		requestedUserUUID uuid.UUID
		ctx               context.Context
	}
)

func (this domain_input) GetRequestedUserUUID() uuid.UUID {

	return this.requestedUserUUID
}

func (this domain_input) GetContext() context.Context {
	return this.ctx
}
