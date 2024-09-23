package assignmentServicePort

import (
	"app/domain/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleAssignnment interface {
		Serve(uuid uuid.UUID, ctx context.Context) (*model.Assignment, error)
		Search(model *model.Assignment, ctx context.Context) (*model.Assignment, error)
	}
)
