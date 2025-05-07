package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"

	"github.com/google/uuid"
)

type (
	CommandGroupUser_CreatedUser_Relation struct {
		aggregateRelation.OneToOneWith[model.User]
		tenantUUID uuid.UUID
	}
)

func (this CommandGroupUser_CreatedUser_Relation) InitializeForeign(join query.IJoinField) {

	join.On("createdBy", "uuid").As("createdUser").
		Filter(
			func(filter query.IFilterExpression) {

				filter.Field("tenantUUID").Equal(this.tenantUUID)
			},
		)
}

func (this CommandGroupUser_CreatedUser_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("createdBy", "uuid").As("createdUser")
	}
}

func (this CommandGroupUser_CreatedUser_Relation) InterceptAfterJoin(queryBuilder query.IQueryBuilder) {

	queryBuilder.Match(
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			return expression.Logical().And(
				func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

					return expression.Filter(
						func(filter query.IDataConditionFilterExpression) {

							filter.Field("assignmentGroups").Not().Equal([]interface{}{})
						},
					)
				},
				func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

					return expression.Logical().Or(
						func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

							return expression.Filter(
								func(filter query.IDataConditionFilterExpression) {

									filter.Field("assignmentGroupMembers").Equal([]interface{}{})
								},
							)
						},
						func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

							return expression.Filter(
								func(filter query.IDataConditionFilterExpression) {

									filter.Field("assignmentGroupMembers.assignmentGroups").Not().Equal([]interface{}{})
								},
							)
						},
					)
				},
			)
		},
	)
}

func (this *CommandGroupUser_CreatedUser_Relation) DispatchDefault(
	tenantUUID uuid.UUID,
) relation.IDBRelationInitiator {

	// return closure.WithForeignFilter(
	// 	*this,
	// 	func(filter query.IFilterExpression) {

	// 		filter.Field("tenantUUID").Equal(tenantUUID)
	// 	},
	// )

	// return closure.WithQueryBuilderInterceptor(
	// 	closure.WithForeignFilter(
	// 		*this,
	// 		func(filter query.IFilterExpression) {

	// 			filter.Field("tenantUUID").Equal(tenantUUID)
	// 		},
	// 	),
	// 	nil,
	// 	func(queryBuilder query.IQueryBuilder) {

	// 		queryBuilder.Match(
	// 			func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 				return expression.Logical().And(
	// 					func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 						return expression.Filter(
	// 							func(filter query.IDataConditionFilterExpression) {

	// 								filter.Field("assignmentGroups").Not().Equal([]interface{}{})
	// 							},
	// 						)
	// 					},
	// 					func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 						return expression.Logical().Or(
	// 							func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 								return expression.Filter(
	// 									func(filter query.IDataConditionFilterExpression) {

	// 										filter.Field("assignmentGroupMembers").Equal([]interface{}{})
	// 									},
	// 								)
	// 							},
	// 							func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 								return expression.Filter(
	// 									func(filter query.IDataConditionFilterExpression) {

	// 										filter.Field("assignmentGroupMembers.assignmentGroups").Not().Equal([]interface{}{})
	// 									},
	// 								)
	// 							},
	// 						)
	// 					},
	// 				)
	// 			},
	// 		)
	// 	},
	// )

	ret := *this

	ret.tenantUUID = tenantUUID

	return ret
}

func (this *CommandGroupUser_CreatedUser_Relation) DispatchForLookupUnAssigned(
	tenantUUID uuid.UUID,
) relation.IDBRelationInitiator {

	return this.DispatchDefault(tenantUUID)
}
