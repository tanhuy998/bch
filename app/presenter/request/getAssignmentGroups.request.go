package requestPresenter

import (
	"app/valueObject/requestInput"

	"github.com/google/uuid"
)

type (
	GetAssignmentGroups struct {
		requestInput.ContextInput
		requestInput.AuthorityInput
		requestInput.TenantMappingInput
		AssignmentUUID *uuid.UUID `param:"uuid" validate:"required"`
	}
)
