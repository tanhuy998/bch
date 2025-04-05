package aggregateRelation

import (
	"app/internal/db/storage"
)

type (
	AbstractRelationInitiator[Repo_Entity_T any] struct {
		Stu storage.IDBStoragePivot[Repo_Entity_T] //mongoStorage.IMongoDBStorageUnit[Repo_Entity_T]
	}
)

func (this *AbstractRelationInitiator[Repo_Entity_T]) GetDBStorageUnitName() string {

	return this.Stu.GetDBStorageUnitName()
}

// func (this *abstract_relation[Repo_Entity_T]) TryJoin(unit interface{}, initializer query.IQueryBuilder) {

// 	switch initiator, accept := unit.(relation.IDBRelationInitiator); {
// 	case accept:
// 		initiator.ResolveRelation(initializer)
// 	default:
// 		return
// 	}
// }

// func (this *abstract_relation[Repo_Entity_T]) PushRelations(
// 	relations ...relation.IDBRelationInitiator,
// ) {

// }
