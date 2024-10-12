package assignmentServicePort

import (
	"app/model"
	"context"

	"github.com/google/uuid"
)

type (
	ICreateAssignment interface {
		Serve(tenantUUID uuid.UUID, data *model.Assignment, ctx context.Context) (*model.Assignment, error)
	}
)
