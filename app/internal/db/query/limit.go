package query

type (
	ISubQueryDataLimit interface {
		Limit(num uint) ISubQueryBuilder
	}
)

type (
	IDataLimit[Model_T any] interface {
		Limit(num uint) IQueryBuilder[Model_T]
	}
)
