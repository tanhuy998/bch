package paginateUseCase

import (
	"app/repository"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ()

type (
	RepoPaginateFunc[Entity_T any, Cursor_T comparable] func(
		collection repository.IMongoRepositoryOperator, cursor Cursor_T, size uint64, ctx context.Context, filters ...primitive.E,
	) ([]Entity_T, error)
)

type (
	PaginateUseCase[Entity_T any] struct {
		//Repo repositoryAPI.IPaginationRepository[Entity_T]
		repo[Entity_T]
	}
)

func (this *PaginateUseCase[Entity_T]) Paginate(
	tenantUUID uuid.UUID, page uint64, size uint64, cursor interface{}, isPrev bool, ctx context.Context,
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
