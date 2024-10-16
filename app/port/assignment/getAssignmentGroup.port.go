package assignmentServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetAssignmentGroups interface {
		Serve(
			tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context,
		) ([]*model.AssignmentGroup, error)
	}
)
