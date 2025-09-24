package mongoRepository

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
	libCommon "app/internal/lib/common"
	repositoryAPI "app/repository/api"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DEFAULT_PAGINATION_SIZE = 10
)

type (
	MongoRepositoryInitOption func(collection *mongo.Collection)
)

type (
	ICRUDMongoRepository[Model_T any] interface {
		IMongoDBRepository
		repositoryAPI.ICRUDRepository[Model_T]
		//storage.IDBStorageUnit[mongo.Collection, Model_T]
		mongoStorage.IMongoDBStorageUnit[Model_T]
	}

	MongoCRUDRepository[Model_T any] struct {
		//mongo_read_projection[Model_T]
		//mongo_filter[Model_T]
		paginate_repository[Model_T]
	}
)

func (this *MongoCRUDRepository[Model_T]) Init(
	db *mongo.Database, collectionName string, options ...MongoRepositoryInitOption,
) *MongoCRUDRepository[Model_T] {

	this.AbstractMongoCollection.Init(db, collectionName)

	collection := this.AbstractMongoCollection.Collection()

	for _, optionFn := range options {

		switch {
		case optionFn == nil:
			continue
		default:
			optionFn(collection)
		}
	}

	return this
}

func (this *MongoCRUDRepository[Model_T]) Clone() repositoryAPI.IPaginationUnit[Model_T] {

	// return libCommon.PointerPrimitive(
	// 	this.mongo_filter,
	// )

	return libCommon.PointerPrimitive(
		*this,
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

	clone := libCommon.PointerPrimitive(
		*this,
	)

	clone.query_condition_materializer.AsFilter(
		fn,
	)

	return clone
}

// func (m *MongoCRUDRepository[Model_T]) Create(model BadType /**tv(Model_T)*/, ctx context.Context) error {
// 	panic("TODO: Remove or impl (available through emb type)")
// }

// func (m *MongoCRUDRepository[Model_T]) UpdateOneByUUID(uuid uuid.UUID, model BadType /**tv(Model_T)*/, ctx context.Context) error {
// 	panic("TODO: Remove or impl (available through emb type)")
// }

//	func (m *MongoCRUDRepository[Model_T]) DeleteMany(model BadType /**tv(Model_T)*/, ctx context.Context) error {
//		panic("TODO: Remove or impl (available through emb type)")
//	}
func (this *MongoCRUDRepository[Model_T]) Match(
	fn repositoryAPI.MatchFunc,
) repositoryAPI.IRepositoryFilterableOperator[Model_T] {

	clone := libCommon.PointerPrimitive(
		*this,
	)

	clone.query_condition_materializer.AsConditionExpression(
		fn,
	)

	return clone
}
