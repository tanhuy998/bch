package paginateUseCase

import (
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PaginationOption func(paginator IPaginatorInitializer)
)

type (
	RepoPaginateFunc[Entity_T any, Cursor_T comparable] func(
		collection repository.IMongoRepositoryOperator, cursor Cursor_T, size uint64, ctx context.Context, filters ...primitive.E,
	) ([]Entity_T, error)
)

type (
	PaginateUseCase[Entity_T any] struct {
		Repo repositoryAPI.IPaginationRepository[Entity_T]
	}
)

func (this *PaginateUseCase[Entity_T]) Paginate(
	tenantUUID uuid.UUID, ctx context.Context, options ...PaginationOption,
) ([]Entity_T, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}

	paginator := NewPaginator(this.Repo)

	for _, fn := range options {

		fn(paginator)
	}

	paginator.IPaginationRepository.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	)

	if paginator.Cursor == nil {

		return paginator.FindOffset(
			paginator.Skip, paginator.Size, ctx,
		)
	}

	if paginator.IsPrev {

		return paginator.FindPrevious(
			"_id", paginator.Cursor, paginator.Size, ctx,
		)
	}

	return paginator.FindNext(
		"_id", paginator.Cursor, paginator.Size, ctx,
	)
}
