package query

import "context"

type (
	IDataRetrievalQuery[Model_T any] interface {
		Find(ctx context.Context) ([]*Model_T, error)
		FindOne(ctx context.Context) (*Model_T, error)
	}

	IDataMutationQuery[Model_T any] interface {
		Update(updateEntity Model_T, ctx context.Context) error
		UpdateOne(updateEntity Model_T, ctx context.Context) error
		Upsert(entity Model_T, ctx context.Context) error
	}

	IDataRemovalQuery[Model_T any] interface {
		Delete(ctx context.Context) error
		DeleteOne(ctx context.Context) error
	}
)
