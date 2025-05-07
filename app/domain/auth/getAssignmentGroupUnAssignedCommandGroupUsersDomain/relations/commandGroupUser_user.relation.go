package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"

	"github.com/google/uuid"
)

type (
	CommandGroupUser_User_Relation struct {
		aggregateRelation.OneToOneWith[model.User]
		tenantUUID uuid.UUID
	}
)

func (this CommandGroupUser_User_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("userUUID", "uuid").As("user").
			Filter(
				func(filter query.IFilterExpression) {
					filter.Field("tenantUUID").Equal(this.tenantUUID)
				},
			)
	}
}

func (this *CommandGroupUser_User_Relation) DispatchDefault(
	tenantUUID uuid.UUID,
) relation.IDBRelationInitiator {

	// return closure.WithForeignFilter(
	// 	this,
	// 	func(filter query.IFilterExpression) {

	// 		filter.Field("tenantUUID").Equal(tenantUUID)
	// 	},
	// )

	clone := *this

	clone.tenantUUID = tenantUUID

	return clone
}

func (this *CommandGroupUser_User_Relation) DispatchForLookupUnAssigned(
	tenantUUID uuid.UUID,
) relation.IDBRelationInitiator {

	return this.DispatchDefault(tenantUUID)
}
