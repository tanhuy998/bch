package join

import "app/internal/db/query"

type (
	join_query_builder struct {
		target_collection string

		filter     query.IFilterGenerator
		projection map[string]uint
		sort       map[string]int
	}
)
