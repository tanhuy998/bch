package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"

	"github.com/google/uuid"
)

type (
	CommandGroupUser_AssignmentGroup_Relation struct {
		aggregateRelation.OneToManyWith[model.AssignmentGroup]
		assignmentGroupUUID uuid.UUID
		tenantUUID          uuid.UUID
	}
)

func (this CommandGroupUser_AssignmentGroup_Relation) InitializeForeign(join query.IJoinField) {

	join.On("commandGroupUUID", "commandGroupUUID").
		As("assignmentGroups").
		Filter(
			func(filter query.IFilterExpression) {

				filter.Field("tenantUUID").Equal(this.tenantUUID)
				filter.Field("uuid").Equal(this.assignmentGroupUUID)
			},
		)
}

func (this CommandGroupUser_AssignmentGroup_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("commandGroupUUID", "commandGroupUUID").As("assignmentGroups")
		// Filter(
		// 	func(filter query.IFilterExpression) {

		// 		filter.Field("tenantUUID").Equal(this.tenantUUID)
		// 		filter.Field("uuid").Equal(this.assignmentGroupUUID)
		// 	},
		// )
	}
}

func (this *CommandGroupUser_AssignmentGroup_Relation) DispatchDefault(
	tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, excludedCommandGroupUserUUIDs []uuid.UUID,
) relation.IDBRelationInitiator {

	// return closure.WithQueryBuilderInterceptor(
	// 	closure.WithForeignFilter(
	// 		this,
	// 		func(filter query.IFilterExpression) {

	// 			filter.Field("uuid").Equal(assignmentGroupUUID)
	// 		},
	// 	),
	// 	func(queryBuilder query.IQueryBuilder) {

	// 		// queryBuilder.Match(
	// 		// 	func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 		// 		return expression.Logical().And(
	// 		// 			func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 		// 				return expression.Filter(
	// 		// 					func(filter query.IDataConditionFilterExpression) {
	// 		// 						filter.Field("tenantUUID").Equal(ctx.GetTenantUUID())
	// 		// 					},
	// 		// 				)
	// 		// 			},
	// 		// 			func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

	// 		// 				return expression.Filter(
	// 		// 					func(filter query.IDataConditionFilterExpression) {

	// 		// 						expectedField := filter.Field("uuid")

	// 		// 						switch len(excludedCommandGroupUserUUIDs) {
	// 		// 						case 0:
	// 		// 							expectedField.Not().Equal(nil)
	// 		// 						default:
	// 		// 							expectedField.Not().EqualOneOfVals(excludedCommandGroupUserUUIDs)
	// 		// 						}
	// 		// 					},
	// 		// 				)
	// 		// 			},
	// 		// 		)
	// 		// 	},
	// 		// )

	// 		queryBuilder.Filter(
	// 			func(filter query.IFilterExpression) {

	// 				filter.Field("tenantUUID").Equal(tenantUUID)

	// 				expectedField := filter.Field("uuid")

	// 				switch len(excludedCommandGroupUserUUIDs) {
	// 				case 0:
	// 					expectedField.Not().Equal(nil)
	// 				default:
	// 					excludedList := make([]interface{}, len(excludedCommandGroupUserUUIDs))

	// 					for i, v := range excludedCommandGroupUserUUIDs {

	// 						excludedList[i] = v
	// 					}

	// 					expectedField.Not().In(excludedList...)
	// 				}
	// 			},
	// 		)
	// 	},
	// 	nil,
	// )

	ret := default_relation_dispatch{
		*this, excludedCommandGroupUserUUIDs,
	}

	ret.tenantUUID = tenantUUID
	ret.assignmentGroupUUID = assignmentGroupUUID

	return ret
}

func (this *CommandGroupUser_AssignmentGroup_Relation) DispatchForLookupUnAssigned(
	tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, lookupCommandGroupUserUUIDs []uuid.UUID,
) relation.IDBRelationInitiator {

	// return closure.WithQueryBuilderInterceptor(
	// 	closure.WithForeignFilter(
	// 		this,
	// 		func(filter query.IFilterExpression) {

	// 			filter.Field("uuid").Equal(assignmentGroupUUID)
	// 		},
	// 	),
	// 	func(queryBuilder query.IQueryBuilder) {

	// 		queryBuilder.Filter(
	// 			func(filter query.IFilterExpression) {

	// 				filter.Field("tenantUUID").Equal(tenantUUID)
	// 				filter.Field("uuid").In(lookupCommandGroupUserUUIDs)
	// 			},
	// 		)
	// 	},
	// 	nil,
	// )

	ret := relation_for_lookup_unassigned{
		*this, lookupCommandGroupUserUUIDs,
	}

	ret.tenantUUID = tenantUUID
	ret.assignmentGroupUUID = assignmentGroupUUID

	return ret
}

type (
	relation_for_lookup_unassigned struct {
		CommandGroupUser_AssignmentGroup_Relation
		lookupCommandGroupUserUUIDs []uuid.UUID
	}
)

func (this relation_for_lookup_unassigned) InterceptBeforeJoin(queryBuilder query.IQueryBuilder) {

	queryBuilder.Filter(
		func(filter query.IFilterExpression) {

			filter.Field("tenantUUID").Equal(this.tenantUUID)
			filter.Field("uuid").In(this.lookupCommandGroupUserUUIDs)
		},
	)
}

type (
	default_relation_dispatch struct {
		CommandGroupUser_AssignmentGroup_Relation
		excludedCommandGroupUserUUIDs []uuid.UUID
	}
)

func (this default_relation_dispatch) InterceptBeforeJoin(queryBuilder query.IQueryBuilder) {

	queryBuilder.Filter(
		func(filter query.IFilterExpression) {

			filter.Field("tenantUUID").Equal(this.tenantUUID)

			expectedField := filter.Field("uuid")

			switch len(this.excludedCommandGroupUserUUIDs) {
			case 0:
				expectedField.Not().Equal(nil)
			default:
				excludedList := make([]interface{}, len(this.excludedCommandGroupUserUUIDs))

				for i, v := range this.excludedCommandGroupUserUUIDs {

					excludedList[i] = v
				}

				expectedField.Not().In(excludedList...)
			}
		},
	)
}
