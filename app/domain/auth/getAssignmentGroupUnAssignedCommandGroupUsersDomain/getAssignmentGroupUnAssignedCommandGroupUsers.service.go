package getAssignmentGroupUnAssignedCommandGroupUsersDomain

import (
	"app/internal/common"
	"app/internal/db/query"
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type (
	GetAssignmentGroupUnAssignedCommandGroupUserService struct {
		AssignmentGroupRepo  repository.IAssignmentGroup
		CommandGroupUserRepo repository.ICommandGroupUser
		Domain_Aggregate
	}
)

// This method is idempotant
func (this *GetAssignmentGroupUnAssignedCommandGroupUserService) Serve(
	// tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, ctx context.Context, exceptCommandGroupUserUUIDs ...uuid.UUID,
	input authServicePort.IGetAssignmentGroupUnAssignedCommandGroupUsersInput, exceptCommandGroupUserUUIDs ...uuid.UUID,
) ([]model.CommandGroupUser, error) {

	switch existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(input.GetAssignmentGroupUUID(), input.GetContext()); {
	case err != nil:
		return nil, err
	case existingAssignmentGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	case *existingAssignmentGroup.TenantUUID != input.GetTenantUUID():
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	// data, err := repository.Aggregate[model.CommandGroupUser](
	// 	this.CommandGroupUserRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 					{
	// 						"uuid", bson.D{
	// 							libCommon.Ternary(
	// 								len(exceptCommandGroupUserUUIDs) == 0,
	// 								bson.E{"$ne", nil},
	// 								bson.E{"$nin", exceptCommandGroupUserUUIDs},
	// 							),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "assignmentGroups"},
	// 					{"localField", "commandGroupUUID"},
	// 					{"foreignField", "commandGroupUUID"},
	// 					{"as", "assignmentGroups"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"uuid", assignmentGroupUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "assignmentGroupMembers"},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "commandGroupUserUUID"},
	// 					{"as", "assignmentGroupMembers"},
	// 					{"pipeline",
	// 						mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 							bson.D{
	// 								{"$lookup",
	// 									bson.D{
	// 										{"from", "assignmentGroups"},
	// 										{"localField", "assignmentGroupUUID"},
	// 										{"foreignField", "uuid"},
	// 										{"as", "assignmentGroups"},
	// 										{
	// 											"pipeline", mongo.Pipeline{
	// 												bson.D{
	// 													{
	// 														"$match", bson.D{
	// 															{"tenantUUID", tenantUUID},
	// 															{
	// 																"assignmentUUID", bson.D{
	// 																	{"$ne", assignmentGroupUUID},
	// 																},
	// 															},
	// 														},
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{
	// 						"$and", bson.A{
	// 							bson.D{
	// 								{
	// 									"assignmentGroups", bson.D{
	// 										{"$ne", bson.A{}},
	// 									},
	// 								},
	// 							},
	// 							bson.D{
	// 								{
	// 									"$or", bson.A{
	// 										bson.D{
	// 											{"assignmentGroupMembers", bson.A{}},
	// 										},
	// 										bson.D{
	// 											{
	// 												"assignmentGroupMembers.assignmentGroups", bson.D{
	// 													{"$ne", bson.A{}},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", "users"},
	// 					{"localField", "userUUID"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "user"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{
	// 						"user", bson.D{
	// 							{"$arrayElemAt", bson.A{"$user", 0}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", "users"},
	// 					{"localField", "createdBy"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "createdUser"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{
	// 						"createdUser", bson.D{
	// 							{"$arrayElemAt", bson.A{"$createdUser", 0}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$project", bson.D{
	// 					{"assignmentGroupMembers", 0},
	// 					{"assignmentGroups", 0},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	data, err := this.Domain_Aggregate.MergeDefaultRelation(
		input.GetTenantUUID(), input.GetAssignmentGroupUUID(), exceptCommandGroupUserUUIDs,
	).Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.ExcludeFields(
				"assignmentGroupMembers", "assignmentGroups",
			)
		},
	).Paginate(input.GetGeneralPaginator(), input.GetContext()) //.All(input.GetContext())

	if err != nil {

		return nil, err
	}

	return data, nil
}

// This method is idempotant
func (this *GetAssignmentGroupUnAssignedCommandGroupUserService) LookupUnAssigned(
	lookupCommandGroupUserUUIDs []uuid.UUID, tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, ctx context.Context,
) ([]model.CommandGroupUser, error) {

	if len(lookupCommandGroupUserUUIDs) == 0 {

		return nil, fmt.Errorf("lookup command group user uuid list is empty")
	}

	switch existingAssignment, err := this.AssignmentGroupRepo.FindOneByUUID(assignmentGroupUUID, ctx); {
	case err != nil:
		return nil, err
	case existingAssignment == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	case *existingAssignment.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	// data, err := repository.Aggregate[model.CommandGroupUser](
	// 	this.CommandGroupUserRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 					{
	// 						"uuid", bson.D{
	// 							{"$in", lookupCommandGroupUserUUIDs},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "assignmentGroups"},
	// 					{"localField", "commandGroupUUID"},
	// 					{"foreignField", "commandGroupUUID"},
	// 					{"as", "assignmentGroups"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"uuid", assignmentGroupUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "assignmentGroupMembers"},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "commandGroupUserUUID"},
	// 					{"as", "assignmentGroupMembers"},
	// 					{"pipeline",
	// 						mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 							bson.D{
	// 								{"$lookup",
	// 									bson.D{
	// 										{"from", "assignmentGroups"},
	// 										{"localField", "assignmentGroupUUID"},
	// 										{"foreignField", "uuid"},
	// 										{"as", "assignmentGroups"},
	// 										{
	// 											"pipeline", mongo.Pipeline{
	// 												bson.D{
	// 													{
	// 														"$match", bson.D{
	// 															{"tenantUUID", tenantUUID},
	// 															{
	// 																"assignmentUUID", bson.D{
	// 																	{"$ne", assignmentGroupUUID},
	// 																},
	// 															},
	// 														},
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{
	// 						"$and", bson.A{
	// 							bson.D{
	// 								{
	// 									"assignmentGroups", bson.D{
	// 										{"$ne", bson.A{}},
	// 									},
	// 								},
	// 							},
	// 							bson.D{
	// 								{
	// 									"$or", bson.A{
	// 										bson.D{
	// 											{"assignmentGroupMembers", bson.A{}},
	// 										},
	// 										bson.D{
	// 											{
	// 												"assignmentGroupMembers.assignmentGroups", bson.D{
	// 													{"$ne", bson.A{}},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", "users"},
	// 					{"localField", "userUUID"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "user"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{
	// 						"user", bson.D{
	// 							{"$arrayElemAt", bson.A{"$user", 0}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", "users"},
	// 					{"localField", "createdBy"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "createdUser"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"tenantUUID", tenantUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{
	// 						"createdUser", bson.D{
	// 							{"$arrayElemAt", bson.A{"$createdUser", 0}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$project", bson.D{
	// 					{"assignmentGroupMembers", 0},
	// 					{"assignmentGroups", 0},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	data, err := this.MergeRelationsForLookupUnAssignedCommandGroupUsers(
		tenantUUID, assignmentGroupUUID, lookupCommandGroupUserUUIDs,
	).Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.ExcludeFields(
				"assignmentGroupMembers", "assignmentGroups",
			)
		},
	).All(ctx)

	if err != nil {

		return nil, err
	}

	return data, nil
}
