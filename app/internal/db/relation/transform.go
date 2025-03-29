package relation

import "app/internal/db/query"

type (
	RelationDataTransformFunc func(transform IRelationLocalDataTransformer)

	ForeignDataExtractFunc func(foreign IRelationForeignField)

	IRelationLocalDataTransformer interface {
		Set(field string) IRelationalDataTransformerSetterExpression
	}

	IRelationalDataTransformerSetterExpression interface {
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
