package relation

import "app/internal/db/query"

type (
	IRelationLocalNavigator interface {
		query.ISubqueryFilterMethod
		query.ISubQueryJoinMethod
		query.ISubQueryProjector
		query.ISubQueryDataLimit
		query.ISkipQueryBuilder
		query.ISubQueryDataSortOrder
		Transform(
			fn RelationDataTransformFunc,
		) IRelationLocalNavigator
	}
)
