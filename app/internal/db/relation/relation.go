package relation

import (
	"app/internal/db/query"
	"app/internal/db/storage"
)

type (
	IDBRelationshipDeclarativeInitializer interface {
		GetRelationInitFunc() query.JoinInitFunc
	}

	IDBRelationshipQueryInitiator interface {
		storage.IDBStorageIdentifier
		IDBRelationshipDeclarativeInitializer
		//GetJoinFieldInitializer() query.IJoinField
		//ResolveRelationQuery(IDBRelationQueryMetadata) Query_T
		ResolveRelation(
			refQueryBuilder IRelationQueryBuilder, foreignInializer IDBRelationNavigator,
		)
	}

	IDBRelationNavigator interface {
		SetLimit(uint64)
		GetLocalField() string
		GetForeignField() string
		GetAliasName() string
	}

	IDBRelationInitiator interface {
		//storage.IDBStorageUnitGetter[DB_Storage_T]
		storage.IDBStorageIdentifier
		//ResolveQuery(fn query.JoinInitFunc) Query_T
		IDBRelationshipQueryInitiator
		IDBRelationKind
	}

	IDBRelationKind interface {
		GetDBRelationKind() string
	}

	IReadRelationQueryBuilder interface {
		query.IQueryBuilder
		// query.ISubqueryFilterMethod
		// query.ISubQueryJoinMethod
		// query.ISubQueryProjector
		// query.ISubQueryDataTransform
		// query.ISubQueryDataLimit
		// query.ISkipQueryBuilder
		// query.ISubQueryDataSortOrder
		PushRelations(relation ...IDBRelationInitiator)
	}

	IClonableReadRelationQueryBuilder interface {
		IReadRelationQueryBuilder
		Clone() IClonableReadRelationQueryBuilder
	}

	IReadRelationQueryBuilderGenerator interface {
		NewRelationQuery() IClonableReadRelationQueryBuilder
	}
)
