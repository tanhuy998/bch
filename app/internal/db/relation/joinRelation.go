package relation

import "app/internal/db/query"

type (
	IDBJoinByRelationQueryBuilder[Source_Repo_Entity_T any, Query_T any, DB_Storage_T any] interface {
		JoinByRelation(
			IDBRelationshipNavigator[Query_T],
		) query.IQueryBuilder[Source_Repo_Entity_T]
	}
)
