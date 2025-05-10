package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	"app/unitOfWork/aggregate/relation/triat/leftJoin"

	"github.com/google/uuid"
)

type (
	Tenant_User_Relation struct {
		leftJoin.OneToManyWith[model.User]
		userUUID uuid.UUID
	}
)

func (this Tenant_User_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return nil
}

func (this Tenant_User_Relation) InitializeForeign(join query.IJoinField) {

	join.On("uuid", "tenantUUID").As("user").
		Match(
			func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

				return expression.Logical().And(
					func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

						return expression.Filter(
							func(filter query.IDataConditionFilterExpression) {

								filter.Field("uuid").Equal(this.userUUID)
							},
						)
					},
					func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

						return expression.Filter(
							func(filter query.IDataConditionFilterExpression) {

								filter.Field("uuid").Not().Equal("$tenantAgent.userUUID")
							},
						)
					},
				)
			},
		)
}

func (this *Tenant_User_Relation) DispatchDefault(
	userUUID uuid.UUID,
) relation.IDBRelationInitiator {

	clone := *this

	clone.userUUID = userUUID

	return clone
}
