package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/relationQueryBuilder/determiner.go"
	"app/internal/db/relation"
	libCommon "app/internal/lib/common"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	NIL_INITIATOR_ERR_MSG = "(mongo relation query builder error) Could not intialize relation_initiator_resovler when the input initiator is nil"
)

type (
	relation_initiator_resolver struct {
		//join_dispatcher queryBuilder.JoinOperationDispatcher
		determiner.RelationJoinDeterminer
		initiator         relation.IDBRelationInitiator
		local_navigator   RelationLocalNavigator
		foreign_navigator RelationForeignNavigator
	}
)

func NewRelationResolver(initiator relation.IDBRelationInitiator) *relation_initiator_resolver {

	if initiator == nil {

		panic(NIL_INITIATOR_ERR_MSG)
	}

	ret := new(relation_initiator_resolver)

	ret.initiator = initiator

	ret.Init()

	return ret
}

func (this *relation_initiator_resolver) Init() {

	if this.initiator == nil {

		panic(NIL_INITIATOR_ERR_MSG)
	}

	this.initLocal()
	this.initForeign()
	this.initJoinDeteminer()
	this.setupJoinOperator()
}

func (this *relation_initiator_resolver) setupJoinOperator() {

	this.local_navigator.PushStages(
		bson.D{
			{"$lookup", this.GetForeignInitializer()},
		},
	)
}

func (this *relation_initiator_resolver) initJoinDeteminer() {

	// this.join_dispatcher.MongoAggregateQueryBuilder = &this.local_navigator.MongoAggregateQueryBuilder
	// this.join_dispatcher.JoinInitializer = &this.foreign_navigator.join_op

	this.RelationJoinDeterminer.SetJoinInitializer(
		&this.foreign_navigator.join_op,
	)

	this.RelationJoinDeterminer.SetLocalQueryBuilder(
		&this.local_navigator.MongoAggregateQueryBuilder,
	)
}

func (this *relation_initiator_resolver) initLocal() {

	this.local_navigator.ptr_foreign_navigator = &this.foreign_navigator
}

func (this *relation_initiator_resolver) initForeign() {

	this.foreign_navigator.join_op.From = this.initiator.GetDBStorageUnitName()
}

func (this *relation_initiator_resolver) GetForeignInitializer() *queryBuilder.JoinOperationInitializer {

	return &this.foreign_navigator.join_op
}

func (this *relation_initiator_resolver) Resolve() {

	this.initiator.ResolveRelation(
		&this.local_navigator,
		&this.foreign_navigator,
	)

	this.foreign_navigator.join_op.Done()

	this.DetermineJoinOperation(this.initiator)

	switch {
	case this.foreign_navigator.unwind_local:

		joinInitializer := this.GetForeignInitializer()

		var unwindPath string = libCommon.Ternary(
			joinInitializer.Alias == "",
			this.initiator.GetDBStorageUnitName(),
			joinInitializer.Alias,
		)

		// switch joinInitializer.Alias {
		// case "":
		// 	unwindPath = fmt.Sprintf(`$%ss`, this.initiator.GetDBStorageUnitName())
		// default:
		// 	unwindPath = fmt.Sprintf("$%s", joinInitializer.Alias)
		// }

		// this.local_navigator.MongoAggregateQueryBuilder.MongoPipeline.PrependStages(
		// 	bson.D{
		// 		{
		// 			"$unwind", bson.D{
		// 				// {
		// 				// 	"path", libCommon.Ternary(
		// 				// 		joinInitializer.Alias == "",
		// 				// 		fmt.Sprintf(`%ss`, this.initiator.GetDBStorageUnitName()),
		// 				// 		joinInitializer.Alias,
		// 				// 	),
		// 				// },
		// 				{"path", unwindPath},
		// 			},
		// 		},
		// 	},
		// )

		this.local_navigator.MongoAggregateQueryBuilder.Unwind(unwindPath)
	}
}
