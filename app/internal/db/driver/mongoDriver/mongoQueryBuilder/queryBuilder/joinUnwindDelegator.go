package queryBuilder

import (
	"app/internal/db/query"
)

type (
	join_unwindable_delegator struct {
		*MongoAggregateQueryBuilder
		foreign_alias string
	}
)

func NewJoinUnwindableDelegator(foreignAlias string, ref *MongoAggregateQueryBuilder) *join_unwindable_delegator {

	ret := &join_unwindable_delegator{
		MongoAggregateQueryBuilder: ref,
		foreign_alias:              foreignAlias,
	}

	return ret
}

func (this *join_unwindable_delegator) UnwindForeign() query.IQueryBuilder {

	this.MongoAggregateQueryBuilder.Unwind(
		this.foreign_alias,
	)

	return this.MongoAggregateQueryBuilder
}
