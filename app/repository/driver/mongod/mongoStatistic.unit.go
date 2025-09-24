package mongoRepository

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
	libError "app/internal/lib/error"
	repositoryAPI "app/repository/api"
	"context"
	"fmt"
)

var (
	default_statistic_error = fmt.Errorf("(mongo statistic unit error)")
)

type (
	mongo_statistic_unit_t[Model_T any] struct {
		mongoStorage.QueryExecutorProxy[Model_T]
		// filter filter.FilterGenerator // mongoRepositoryFilter.MongoRepositoryFilterGenerator
		query_condition_materializer filter_condition_materializer_t
	}
)

func (this mongo_statistic_unit_t[Model_T]) Clone() *mongo_statistic_unit_t[Model_T] {

	// this.filter = nil

	this.query_condition_materializer.f = nil

	return &this
}

func (this mongo_statistic_unit_t[Model_T]) Count(
	baseCtx context.Context,
) (affectedDocsCount int64, err error) {

	// ret, err := this.QueryExecutorProxy.CountDocuments(baseCtx, this.filter.Get())

	filter, err := this.query_condition_materializer.getFilterExpression()

	if err != nil {

		return -1, err
	}

	ret, err := this.QueryExecutorProxy.CountDocuments(baseCtx, filter)

	if err != nil {

		return -1, libError.NewInternal(default_statistic_error, err)
	}

	return ret, nil
}

func (this *mongo_statistic_unit_t[Model_T]) SelfStatistic() repositoryAPI.IStatisticAffectedCountableUnit {

	return this
}

func (this *mongo_statistic_unit_t[Model_T]) Statistic(fn repositoryAPI.FilterFunc) repositoryAPI.IStatisticUnit {

	clone := this.Clone()

	clone.query_condition_materializer.f = nil

	clone.query_condition_materializer.AsFilter(fn)

	return clone
}
