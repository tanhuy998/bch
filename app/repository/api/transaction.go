package repositoryAPI

import "context"

type (
	ITransactionDBClient interface {
		WithTransaction(ctx context.Context, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
	}
)
