package mongoQueryBuilder

import "app/internal/db/query"

type (
	join_op_pipeline struct {
		ref_join_op *JoinOperationInitializer
		*MongoAggregateQueryBuilder
	}
)

func NewJoinOperationPipeline(ref *JoinOperationInitializer) *join_op_pipeline {

	return &join_op_pipeline{
		ref_join_op: ref,
	}
}

func (this *join_op_pipeline) Limit(number uint64) query.IQueryBuilder {

	if number <= 1 {

		this.ref_join_op.is_unwind = true

	} else {

		this.ref_join_op.is_unwind = false
	}

	this.MongoAggregateQueryBuilder.Limit(number)

	return this
}
