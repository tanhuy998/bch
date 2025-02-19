package mongoRelation

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
)

type (
	asbtract_relation[Repo_Entity_T any] struct {
		mongoStorage.IMongoDBStorageUnit[Repo_Entity_T]
	}
)

// func (this *asbtract_relation) GetJoinFieldInitializer() query.IJoinField {

// 	return new(mongoQueryBuilder.JoinOperationInitializer)
// }
