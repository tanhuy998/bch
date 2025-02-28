package relation

import (
	"app/internal/db/query"
	"app/internal/db/storage"
)

type (
	IDBRelationshipDeclarativeInitializer interface {
		GetRelationInitFunc() query.JoinInitFunc
	}

	IDBRelationshipQueryInitializer[Query_Meta_T any] interface {
		storage.IDBStorageIdentifier
		IDBRelationshipDeclarativeInitializer
		//GetJoinFieldInitializer() query.IJoinField
		ResolveRelationQuery(IDBRelationQueryMetadata) Query_Meta_T
	}

	IDBRelationQueryMetadata interface {
		SetLimit(uint64)
		GetLocalField() string
		GetForeignField() string
		GetAliasName() string
	}

	IDBRelationshipNavigator[Query_T any] interface {
		//storage.IDBStorageUnitGetter[DB_Storage_T]
		storage.IDBStorageIdentifier
		//ResolveQuery(fn query.JoinInitFunc) Query_T
		IDBRelationshipQueryInitializer[Query_T]
	}
)
