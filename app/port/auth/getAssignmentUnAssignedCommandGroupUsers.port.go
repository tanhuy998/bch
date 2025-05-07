package authServicePort

import (
	contextHolderPort "app/port/contextHolder"
	paginateServicePort "app/port/paginate"
	"app/unitOfWork/aggregate"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentGroupUnAssignedCommandGroupUsersInput interface {
		//context.Context
		aggregate.IDomainContext
		aggregate.IAssignmentGroupDomain
		contextHolderPort.IContextHolder
		paginateServicePort.IGeneralPaginator
	}
)

type (
	IGetAssignmentGroupUnAssignedCommandGroupUsers[Entity_T any] interface {
		Serve(
			//TenantUUID, AssignmentGroupUUID uuid.UUID, ctx context.Context, exceptCommandGroupUUIDs ...uuid.UUID,
			input IGetAssignmentGroupUnAssignedCommandGroupUsersInput, excludedUserUUIDs ...uuid.UUID,
		) ([]Entity_T, error)
		LookupUnAssigned(
			lookupCommandGroupUUIDs []uuid.UUID, tenantUUID uuid.UUID, AssignmentGroupUUID uuid.UUID, ctx context.Context,
		) ([]Entity_T, error)
	}
)
