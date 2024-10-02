package crudResourcePort

import (
	"context"
)

type (
	ICreateResource[T any] interface {
		CreateByModel(model *T, ctx context.Context) (*T, error)
	}

	ISearchResouce[T any] interface {
		SearchByModel(mode *T, ctx context.Context) (*T, error)
	}
)
