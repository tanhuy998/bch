package assignmentServicePort

import (
	"app/domain/model"
	"context"
)

type (
	IGetSingleAssignnment interface {
		Serve(uuid string, ctx context.Context) (*model.Assignment, error)
		Search(model *model.Assignment, ctx context.Context) (*model.Assignment, error)
	}
)
