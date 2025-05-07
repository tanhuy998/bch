package getUserParticipatedCommandGroupsDomain

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	"app/unitOfWork/aggregate"
	aggregateRelation "app/unitOfWork/aggregate/relation"

	"github.com/google/uuid"
)

type (
	InputContext interface {
		aggregate.IDomainContext
		aggregate.IDomainUser
	}
)

type (
	CommandGroup_CommandGroupUser_Relation struct {
		aggregateRelation.OneToOneWith[model.CommandGroupUser]
		tenantUUID uuid.UUID
		userUUID   uuid.UUID
	}
)

func (this CommandGroup_CommandGroupUser_Relation) InitializeForeign(join query.IJoinField) {

	join.On("uuid", "commandGroupUUID").As("commandGroupUser").
		Filter(
			func(filter query.IFilterExpression) {

				filter.Field("tenantUUID").Equal(this.tenantUUID)
				filter.Field("userUUID").Equal(this.userUUID)
			},
		)
}

func (this CommandGroup_CommandGroupUser_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("uuid", "commandGroupUUID").As("commandGroupUser")
	}
}

func (this CommandGroup_CommandGroupUser_Relation) InterceptBeforeJoin(queryBuilder query.IQueryBuilder) {

	queryBuilder.Filter(
		func(filter query.IFilterExpression) {

			filter.Field("tenantUUID").Equal(this.tenantUUID)
		},
	)
}

func (this *CommandGroup_CommandGroupUser_Relation) DispatchDefault(
	tenantUUID uuid.UUID, userUUID uuid.UUID,
) relation.IDBRelationInitiator {

	// return closure.WithQueryBuilderInterceptor(
	// 	closure.WithForeignFilter(
	// 		this,
	// 		func(filter query.IFilterExpression) {

	// 			filter.Field("tenantUUID").Equal(tenantUUID)
	// 			filter.Field("userUUID").Equal(userUUID)
	// 		},
	// 	),
	// 	func(queryBuilder query.IQueryBuilder) {

	// 		queryBuilder.Filter(
	// 			func(filter query.IFilterExpression) {

	// 				filter.Field("tenantUUID").Equal(tenantUUID)
	// 			},
	// 		)
	// 	},
	// 	nil,
	// )

	clone := *this

	clone.tenantUUID = tenantUUID
	clone.userUUID = userUUID

	return clone
}
