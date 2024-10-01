package assignmentServicePort

import (
	"app/src/model"
	"context"

	"github.com/google/uuid"
)

type (
	IGetSingleAssignmentGroup interface {
		Serve(uuid uuid.UUID, ctx context.Context) (*model.AssignmentGroup, error)
		// SearchByModel(data *model.AssignmentGroup, ctx context.Context) (*model.AssignmentGroup, error)
	}
)
