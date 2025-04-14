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
		//requestInput.RangePaginateInput
		requestInput.PaginateInput
		AssignmentUUID *uuid.UUID `param:"uuid" validate:"required"`
	}
)

func (this *GetAssignmentGroups) GetAssignmentUUID() uuid.UUID {

	return *this.AssignmentUUID
}
