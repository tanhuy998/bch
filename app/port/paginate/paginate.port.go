package paginateServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IPaginate[Entity_T any, Cursor_T comparable] interface {
		Paginate(
			tenantUUID uuid.UUID, page uint64, size uint64, cursor *Cursor_T, isPrev bool, ctx context.Context,
		) ([]Entity_T, error)
	}
)
