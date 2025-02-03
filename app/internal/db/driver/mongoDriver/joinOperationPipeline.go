package mongoDriver

import "app/internal/db/query"

type (
	join_op_pipeline struct {
		ref_join_op *join_op
		*mongo_query
	}
)

func NewJoinOperationPipeline(ref *join_op) *join_op_pipeline {

	return &join_op_pipeline{
		ref_join_op: ref,
	}
}

func (this *join_op_pipeline) Limit(number uint) query.ISubQueryBuilder {

	if number <= 1 {

		this.ref_join_op.is_unwind = true

	} else {

		this.ref_join_op.is_unwind = false
	}

	this.mongo_query.Limit(number)

	return this
}
