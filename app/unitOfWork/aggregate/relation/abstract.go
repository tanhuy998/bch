package aggregateRelation

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
)

type (
	abstract_relation[Repo_Entity_T any] struct {
		Stu mongoStorage.IMongoDBStorageUnit[Repo_Entity_T]
	}
)

func (this *abstract_relation[Repo_Entity_T]) GetDBStorageUnitName() string {

	return this.Stu.GetDBStorageUnitName()
}

// func (this *abstract_relation[Repo_Entity_T]) TryJoin(unit interface{}, initializer query.IQueryBuilder) {

// 	switch initiator, accept := unit.(relation.IDBRelationInitiator); {
// 	case accept:>
// 		initiator.ResolveRelation(initializer)
// 	default:
// 		return
// 	}
// }
