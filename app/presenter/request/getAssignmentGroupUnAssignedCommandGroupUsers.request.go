package requestPresenter

import (
	"app/valueObject/requestInput"

	"github.com/google/uuid"
)

type (
	GetAssignmentGroupUnAssignedCommandGroupUsers struct {
		requestInput.PaginateInput
		requestInput.ContextInput
		requestInput.TenantMappingInput
		requestInput.AuthorityInput
		AssignmentGroupUUID *uuid.UUID `param:"assignmentGroupUUID" validate:"required"`
	}
)
