package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
	libCommon "app/internal/lib/common"
)

type (
	joinable_collection[Model_T, Target_Entity_T any] struct {
		mongoStorage.MongoDBQueryMonitorCollection
	}
)

func (this *joinable_collection[Model_T, Target_Entity_T]) Clone() *joinable_collection[Model_T, Target_Entity_T] {

	return libCommon.PointerPrimitive(*this)
}

// func (this *joinable_collection[Model_T, Target_Entity_T]) Join(
// 	another string, fn func(query.IJoinField),
// ) query.IQueryBuilder[Target_Entity_T] {

// 	return join.NewJoinQueryBuilder()
// }
