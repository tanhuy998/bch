package authServicePort

import (
	"app/model"
	paginateServicePort "app/port/paginate"
	"context"

	"github.com/google/uuid"
)

type (
	IGetCommandGroupUsers[CusorType comparable] interface {
		Serve(tenantUUID uuid.UUID, groupUUID uuid.UUID, paginator paginateServicePort.IPaginator[CusorType], ctx context.Context) ([]model.CommandGroupUser, error)
	}
)
