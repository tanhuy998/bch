package query

type (
	ISubQueryDataLimit interface {
		Limit(num uint64) IQueryBuilder
	}
)

type (
	IDataLimit[Model_T any] interface {
		Limit(num uint64) IGenericQueryBuilder[Model_T]
	}
)
