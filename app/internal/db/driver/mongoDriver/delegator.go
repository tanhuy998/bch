package mongoDriver

import (
	"app/internal/db"
	"app/internal/db/query"
	libCommon "app/internal/lib/common"
)

type (
	MongoDBDelegator[Entity_T any, Source_Repo_Entity_T any] struct {
		//MongoDBQueryMonitorCollection
		Su db.IDBStorageUnit[MongoDBQueryMonitorCollection, Source_Repo_Entity_T]
	}
)

func (this *MongoDBDelegator[Entity_T, Source_Repo_Entity_T]) Clone() *MongoDBDelegator[Entity_T, Source_Repo_Entity_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *MongoDBDelegator[Entity_T, Source_Repo_Entity_T]) Query() query.IQueryBuilder[Entity_T] {

	return NewDelegatorQueryBuilder[Entity_T](this.Su.GetStorageUnit())
}

func (this *MongoDBDelegator[Entity_T, Source_Repo_Entity_T]) Exec(query query.IQueryBuilder[Entity_T]) query.IQueryExecutor[Entity_T] {

	executor := query.NewExecutor(this)

	return executor
}
