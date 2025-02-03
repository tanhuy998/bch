package query

type (
	DataTransformFunc = func(IDataTransformer)

	IDataTransformer interface {
		Set(field string) IDataTransformSetter
	}

	IDataTransformSetter interface {
		Value(val interface{})
		Ref(field string)
	}
)

type (
	ISubQueryDataTransform interface {
		Transform(DataTransformFunc) ISubQueryBuilder
	}
)

type (
	IDataTransform[Model_T any] interface {
		Transform(DataTransformFunc) IQueryBuilder[Model_T]
	}
)
