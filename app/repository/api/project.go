package repositoryAPI

import "context"

type (
	IRepositoryProjectableOperator[Model_T any] interface {
		Find(ctx context.Context) ([]*Model_T, error)
		FindOne(ctx context.Context) (*Model_T, error)
	}

	IProjector[Model_T any] interface {
		Select(fields ...string) IRepositoryProjectableOperator[Model_T]
		ExcludeFields(fields ...string) IRepositoryProjectableOperator[Model_T]
	}
)
