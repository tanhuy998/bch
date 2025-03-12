package query

type (
	ISkipQueryBuilder interface {
		Skip(uint64) IQueryBuilder
	}
)
