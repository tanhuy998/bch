package repositoryAPI

import (
	"context"
)

type (
	// IFilterRepository[Model_T any] interface {
	// 	Filter(filter interface{}) ICRUDRepository[Model_T]
	// }

	IRepositoryFilterableOperator[Model_T any] interface {
		//IRepositoryReadOperator[Model_T]
		IRepositoryProjectableOperator[Model_T]
		IProjectionMethods[Model_T]
		Find(ctx context.Context) ([]*Model_T, error)
		FindOne(ctx context.Context) (*Model_T, error)
		Update(updateEntity Model_T, ctx context.Context) error
		UpdateOne(updateEntity Model_T, ctx context.Context) error
		Delete(ctx context.Context) error
		DeleteOne(ctx context.Context) error
		Upsert(entity Model_T, ctx context.Context) error
	}

	IFilterGenerator interface {
		Add(...interface{}) IFilterGenerator
		IFilterExpression
	}

	IFilterExpression interface {
		Field(name string) IFilterExpressionOperator
	}

	IFilterLogicalOperator interface {
		Or(FilterLogicalGroupFunc)
		And(FilterLogicalGroupFunc)
	}

	IFilterExpressionOperator interface {
		IComaparisonOperator
		Not() IComaparisonOperator
	}

	IComaparisonOperator interface {
		//IFilterLogicalOperator
		Equal(val interface{})
		GreaterThan(val interface{})
		GreaterOrEqual(val interface{})
		LessThan(val interface{})
		LessThanOrEqual(val interface{})
	}

	FilterLogicalGroupFunc = func(filteredField IFilterExpressionOperator)
	FilterFunc             = func(filter IFilterGenerator)

	IFilterMethods[Model_T any] interface {
		Filter(FilterFunc) IRepositoryFilterableOperator[Model_T]
	}
)
