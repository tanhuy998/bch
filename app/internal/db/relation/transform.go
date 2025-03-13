package relation

import "app/internal/db/query"

type (
	RelationDataTransformFunc func(transform IRelationDataTransformer)

	ForeignDataExtractFunc func(foreign IRelationForeignField)

	IRelationQueryBuilder interface {
		query.ISubqueryFilterMethod
		query.ISubQueryJoinMethod
		query.ISubQueryProjector
		query.ISubQueryDataLimit
		query.ISkipQueryBuilder
		query.ISubQueryDataSortOrder
		Transform(
			fn RelationDataTransformFunc,
		)
	}

	IRelationDataTransformer interface {
		Set(field string) IRelationDataTransformerSetterExpression
	}

	IRelationDataTransformerSetterExpression interface {
		query.IDataTransformSetter
		AsForeign() IRelationForeignField
	}

	IRelationForeignFieldListType interface {
		FirstElement()
		At(index uint64)
		Count()
	}

	IRelationForeignField interface {
		IRelationForeignFieldListType
	}
)
