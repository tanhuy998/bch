package repositoryAPI

import "app/internal/db/query"

type (
	MatchFunc query.DataConditionMatchFunc

	IConditionMethod[Model_T any] interface {
		Match(fn MatchFunc) IRepositoryFilterableOperator[Model_T]
	}
)
