package mongoRelation

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
)

type (
	abstract_relation[Repo_Entity_T any] struct {
		Stu mongoStorage.IMongoDBStorageUnit[Repo_Entity_T]
	}
)

// func (this *asbtract_relation) GetJoinFieldInitializer() query.IJoinField {

// 	return new(mongoQueryBuilder.JoinOperationInitializer)
// }

func (this *abstract_relation[Repo_Entity_T]) GetDBStorageUnitName() string {

	return this.Stu.GetDBStorageUnitName()
}
