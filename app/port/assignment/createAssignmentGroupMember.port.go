package assignmentServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	ICreateAssignmentGroupMember interface {
		Serve(
			tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, commandGroupUserUUIDList []*model.AssignmentGroupMember, ctx context.Context,
		) ([]*model.AssignmentGroupMember, error)
	}
)
