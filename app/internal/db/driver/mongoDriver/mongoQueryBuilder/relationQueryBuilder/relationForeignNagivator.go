package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"
	libCommon "app/internal/lib/common"
)

type (
	/*
		RelationForeignNavigator is the wrapper of a join operation in mongodb,
		it's called navigator because it plays a role for the relation initiator
		to indentify and manipulate the relation.
	*/
	RelationForeignNavigator struct {
		join_op      queryBuilder.JoinOperationInitializer
		unwind_local bool
		//local_unwind_delegator query.IUnwindJoinedData
	}
)

func NewRelationForeignNavigator() *RelationForeignNavigator {

	ret := new(RelationForeignNavigator)

	ret.join_op.JoinPipeline.Init()

	return ret
}

func (this *RelationForeignNavigator) SetLimit(num uint64) {

	num = libCommon.Ternary(num == 0, 1, num)

	this.join_op.JoinPipeline.Limit(num)
}

func (this *RelationForeignNavigator) GetLocalField() string {

	return this.join_op.Local_field
}

func (this *RelationForeignNavigator) GetForeignField() string {

	return this.join_op.Foreign_field
}

func (this *RelationForeignNavigator) GetAliasName() string {

	alias := this.join_op.Alias

	switch alias {
	case "":
		return this.join_op.From
	default:
		return alias
	}
}

func (this *RelationForeignNavigator) Manipulate(fn relation.ForeignManipulatorFunc) {

	if fn == nil {

		panic("foreign manipulator func must not be nil")
	}

	this.join_op.JoinPipeline.Init()

	dispatcher := NewRelationDispatcher(
		&this.join_op.JoinPipeline.MongoAggregateQueryBuilder,
	)

	fn(dispatcher)

	this.join_op.Done()
}

func (this *RelationForeignNavigator) UnwindLocal() {

	this.unwind_local = true
}
