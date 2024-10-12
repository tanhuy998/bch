package assignmentServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IModifyAssignment interface {
		Serve(
			tenantUUID uuid.UUID, assignmentUUID uuid.UUID, dataModel *model.Assignment, ctx context.Context,
		) error
	}
)
