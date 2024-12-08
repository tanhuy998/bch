package genericUseCase

import (
	"app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	RepoPaginateFunc[Entity_T any, Cursor_T comparable] func(
		collection repository.IMongoRepositoryOperator, cursor Cursor_T, size uint64, ctx context.Context, filters ...primitive.E,
	) ([]Entity_T, error)
)

type (
	PaginateUseCase[
		Entity_T any,
		Cursor_T comparable,
		Filter_T, Projection_T any,
		Repository_T repository.IPaginateRepository[Filter_T, Projection_T, Entity_T, Cursor_T],
	] struct {
		Repo Repository_T
	}
)

func (this *PaginateUseCase[Entity_T, Cursor_T, Filter_T, Projection_T, Repository_T]) Paginate(
	tenantUUID uuid.UUID, page uint64, size uint64, cursor *Cursor_T, isPrev bool, ctx context.Context, filters Filter_T,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	var f bson.D = bson.D{
		{"tenantUUID", tenantUUID},
	}

	if len(filters) > 0 {

		f = append(f, filters...)
	}

	if cursor == nil {

		return this.Repo.FindOffset(
			f,
			page,
			size,
			nil,
			ctx,
		)
	}

	if isPrev {

		return this.Repo.FindPrevious(
			*cursor, size, ctx, f,
		)
	}

	return this.Repo.FindNext(
		*cursor, size, ctx, f,
	)
}
