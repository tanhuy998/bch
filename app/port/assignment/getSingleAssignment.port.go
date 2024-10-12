package assignmentServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleAssignnment interface {
		Serve(tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context) (*model.Assignment, error)
		Search(model *model.Assignment, ctx context.Context) (*model.Assignment, error)
	}
)
