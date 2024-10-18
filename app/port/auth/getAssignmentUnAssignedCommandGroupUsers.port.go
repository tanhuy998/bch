package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentGroupUnAssignedCommandGroupUsers interface {
		Serve(
			TenantUUID, AssignmentGroupUUID uuid.UUID, ctx context.Context, exceptCommandGroupUUIDs ...uuid.UUID,
		) ([]*model.CommandGroupUser, error)
		LookupUnAssigned(
			lookupCommandGroupUUIDs []uuid.UUID, tenantUUID, AssignmentGroupUUID uuid.UUID, ctx context.Context,
		) ([]*model.CommandGroupUser, error)
	}
)
