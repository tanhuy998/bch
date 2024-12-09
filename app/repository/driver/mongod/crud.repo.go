package mongoRepository

import (
	libCommon "app/internal/lib/common"
	repositoryAPI "app/repository/api"
)

const (
	DEFAULT_PAGINATION_SIZE = 10
)

type (
	ICRUDMongoRepository[Model_T any] interface {
		IMongoDBRepository
		repositoryAPI.ICRUDRepository[Model_T]
	}

	MongoCRUDRepository[Model_T any] struct {
		//mongo_read_projection[Model_T]
		mongo_filter[Model_T]
	}
)

func (this *MongoCRUDRepository[Model_T]) clone() *MongoCRUDRepository[Model_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *MongoCRUDRepository[Model_T]) Select(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	return this.clone().Select(fields...)
}

func (this *MongoCRUDRepository[Model_T]) ExcludeFields(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	return this.clone().ExcludeFields(fields...)
}

func (this *MongoCRUDRepository[Model_T]) Filter(
	fn repositoryAPI.FilterFunc,
) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

	return this.clone().Filter(fn)
}
