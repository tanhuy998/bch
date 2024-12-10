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

func (this *MongoCRUDRepository[Model_T]) Clone() *mongo_filter[Model_T] {

	return libCommon.PointerPrimitive(
		this.mongo_filter,
	)
}

func (this *MongoCRUDRepository[Model_T]) Select(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	return this.Clone().Select(fields...)
}

func (this *MongoCRUDRepository[Model_T]) ExcludeFields(fields ...string) (ret repositoryAPI.IRepositoryProjectableOperator[Model_T]) {

	return this.Clone().ExcludeFields(fields...)
}

func (this *MongoCRUDRepository[Model_T]) Filter(
	fn repositoryAPI.FilterFunc,
) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

	return this.Clone().Filter(fn)
}

// func (m *MongoCRUDRepository[Model_T]) Create(model BadType /**tv(Model_T)*/, ctx context.Context) error {
// 	panic("TODO: Remove or impl (available through emb type)")
// }

// func (m *MongoCRUDRepository[Model_T]) UpdateOneByUUID(uuid uuid.UUID, model BadType /**tv(Model_T)*/, ctx context.Context) error {
// 	panic("TODO: Remove or impl (available through emb type)")
// }

// func (m *MongoCRUDRepository[Model_T]) DeleteMany(model BadType /**tv(Model_T)*/, ctx context.Context) error {
// 	panic("TODO: Remove or impl (available through emb type)")
// }
