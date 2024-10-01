package assignmentServicePort

import (
	"app/src/model"
	"context"
)

type (
	ICreateAssignment interface {
		Serve(data *model.Assignment, ctx context.Context) (*model.Assignment, error)
	}
)
