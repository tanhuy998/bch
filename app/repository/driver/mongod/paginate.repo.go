package mongoRepository

import (
	repositoryAPI "app/repository/api"
)

type (
	paginate_repository[Model_T any] struct {
		mongo_filter[Model_T]
	}
)

func (this *paginate_repository[Model_T]) QueryConditionUnit() repositoryAPI.IPaginateQueryConditionUnit {

	return &this.query_condition_materializer
}
