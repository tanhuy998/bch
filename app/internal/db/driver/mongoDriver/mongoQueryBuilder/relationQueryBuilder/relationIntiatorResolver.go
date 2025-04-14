package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"
	libCommon "app/internal/lib/common"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	NIL_INITIATOR_ERR_MSG = "(mongo relation query builder error) Could not intialize relation_initiator_resovler when the input initiator is nil"
)

type (
	relation_initiator_resolver struct {
		initiator         relation.IDBRelationshipQueryInitiator
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

	switch {
	case this.foreign_navigator.unwind_local:

		joinInitializer := this.GetForeignInitializer()

		this.local_navigator.MongoAggregateQueryBuilder.MongoPipeline.PrependStages(
			bson.D{
				{
					"$unwind", bson.D{
						{
							"path", libCommon.Ternary(
								joinInitializer.Alias == "",
								fmt.Sprintf(`%ss`, this.initiator.GetDBStorageUnitName()),
								joinInitializer.Alias,
							),
						},
					},
				},
			},
		)
	}
}
