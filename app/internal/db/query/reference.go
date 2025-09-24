package query

type (
	IQuerySelfReference interface {
		Origin() string
		GetQuerySelfReference() string
	}
)
