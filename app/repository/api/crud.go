package repositoryAPI

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IFindMany[Model_T any] interface {
		FindMany(query bson.D, ctx context.Context, projection ...bson.E) ([]*Model_T, error)
	}

	ICreateMany[Model_T any] interface {
		CreateMany(models []*Model_T, ctx context.Context) error
	}

	IFindByFilterRepository[Filter_T, Projection_T, Model_T any] interface {
		Find(filter Filter_T, ctx context.Context) (*Model_T, error)

		FindMany(filter Filter_T, ctx context.Context, projection ...Projection_T) ([]*Model_T, error)
	}

	IFindByOffsetRepository[Filter_T, Projection_T, Model_T any] interface {
		FindOffset(
			filter Filter_T, offset uint64, size uint64, sort *bson.D, ctx context.Context, projections ...Projection_T,
		) ([]Model_T, error)
	}

	IUpdateByFilterRepository[Model_T any] interface {
		UpdateManyByFilter(filter interface{}, update bson.D, ctx context.Context) error
		UpsertManyByFilter(filter interface{}, update bson.D, ctx context.Context) error
	}

	IDeleteByFilterRepository[Model_T any] interface {
		DeleteManyByFilter(filter interface{}, ctx context.Context) error
	}

	// IPaginateRepository[Filter_T, Projection_T, Model_T any, Cursor_T comparable] interface {
	// 	IRepositoryReadOperator[Model_T]
	// 	IFindByOffsetRepository[Filter_T, Projection_T, Model_T]
	// 	FindNext(cursor Cursor_T, size uint64, ctx context.Context, filters Filter_T) ([]Model_T, error)
	// 	FindPrevious(cursor Cursor_T, size uint64, ctx context.Context, filters Filter_T) ([]Model_T, error)
	// }

	ICRUDRepository[Model_T any] interface {
		ICreateMany[Model_T]
		IRepositoryReadOperator[Model_T]
		IRepositoryFilterableOperator[Model_T]
		IFilterMethods[Model_T]
		IProjectionMethods[Model_T]
		Create(model *Model_T, ctx context.Context) error

		// FindOffset(
		// 	query interface{}, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
		// ) ([]Model_T, error)
		//reate(model *Model_T, ctx context.Context) error

		// FindOffset(
		// 	query interface{}, offset uint64, size uint64, sort *bson.D, ctx context.Context, projection ...bson.E,
		// ) ([]Model_T, error)
		UpdateOneByUUID(uuid uuid.UUID, model *Model_T, ctx context.Context) error
		DeleteMany(model *Model_T, ctx context.Context) error
	}

	ICRUDMongoRepository[Model_T any] interface {
		ICRUDRepository[Model_T]
		IMongoDBRepository
	}
)
