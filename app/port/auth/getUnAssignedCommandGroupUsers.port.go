package authServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetUnAssignedCommandGroupUsers interface {
		Serve(
			TenantUUID, AssignmentUUID uuid.UUID, ctx context.Context, exceptCommandGroupUUIDs ...uuid.UUID,
		) ([]*model.CommandGroupUser, error)
		LookupUnAssigned(
			lookupCommandGroupUUIDs []uuid.UUID, tenantUUID, AssignmentUUID uuid.UUID, ctx context.Context,
		) ([]*model.CommandGroupUser, error)
	}
)
