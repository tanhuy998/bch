package paginateOutput

type (
	ignore_optional_query_t struct{}
)

type (
	OptionalQuerySetFunc[Optional_Query_T any] func(ref *Optional_Query_T)

	IOptionalQueryModifier[T any] interface {
		SetOptionalQuery(fn OptionalQuerySetFunc[T])
	}
)
