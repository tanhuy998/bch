package getUnAssignedCommandGroupUsersDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetUnAssignedCommandGroupUser struct {
		AssignmentRepo       repository.IAssignment
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

// This method is idempotant
func (this *GetUnAssignedCommandGroupUser) Serve(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context, exceptCommandGroupUserUUIDs ...uuid.UUID,
) ([]*model.CommandGroupUser, error) {

	switch existingAssignment, err := this.AssignmentRepo.FindOneByUUID(assignmentUUID, ctx); {
	case err != nil:
		return nil, err
	case existingAssignment == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignmen not found"))
	case *existingAssignment.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	data, err := repository.Aggregate[model.CommandGroupUser](

		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"tenantUUID", tenantUUID},
						{
							"uuid", bson.D{
								{"$nin", exceptCommandGroupUserUUIDs},
							},
						},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "assignmentGroupMembers"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUserUUID"},
						{"as", "assignmentGroupMembers"},
						{"pipeline",
							mongo.Pipeline{
								bson.D{
									{"$lookup",
										bson.D{
											{"from", "assignmentGroups"},
											{"localField", "assignmentGroupUUID"},
											{"foreignField", "uuid"},
											{"as", "assignmentGroups"},
											{
												"pipeline", mongo.Pipeline{
													bson.D{
														{
															"$match", bson.D{
																{"tenantUUID", tenantUUID},
																{
																	"assignmentUUID", bson.D{
																		{"$ne", assignmentUUID},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								bson.D{
									{
										"$match", bson.D{
											{"tenantUUID", tenantUUID},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{
							"$or", bson.A{
								bson.D{
									{"assignmentGroupMembers", bson.A{}},
								},
								bson.D{
									{
										"assignmentGroupMembers.assignmentGroups", bson.D{
											{"$ne", bson.A{}},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{
					"$project", bson.D{
						{"assignmentGroupMembers", 0},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return data, nil
}

// This method is idempotant
func (this *GetUnAssignedCommandGroupUser) LookupUnAssigned(
	lookupCommandGroupUserUUIDs []uuid.UUID, tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context,
) ([]*model.CommandGroupUser, error) {

	switch existingAssignment, err := this.AssignmentRepo.FindOneByUUID(assignmentUUID, ctx); {
	case err != nil:
		return nil, err
	case existingAssignment == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	case *existingAssignment.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	data, err := repository.Aggregate[model.CommandGroupUser](

		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"tenantUUID", tenantUUID},
						{
							"uuid", bson.D{
								{"$in", lookupCommandGroupUserUUIDs},
							},
						},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "assignmentGroupMembers"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUserUUID"},
						{"as", "assignmentGroupMembers"},
						{"pipeline",
							mongo.Pipeline{
								bson.D{
									{"$lookup",
										bson.D{
											{"from", "assignmentGroups"},
											{"localField", "assignmentGroupUUID"},
											{"foreignField", "uuid"},
											{"as", "assignmentGroups"},
											{
												"pipeline", mongo.Pipeline{
													bson.D{
														{
															"$match", bson.D{
																{"tenantUUID", tenantUUID},
																{
																	"assignmentUUID", bson.D{
																		{"$ne", assignmentUUID},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
								bson.D{
									{
										"$match", bson.D{
											{"tenantUUID", tenantUUID},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{
							"$or", bson.A{
								bson.D{
									{"assignmentGroupMembers", bson.A{}},
								},
								bson.D{
									{
										"assignmentGroupMembers.assignmentGroups", bson.D{
											{"$ne", bson.A{}},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{
					"$project", bson.D{
						{"assignmentGroupMembers", 0},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return data, nil
}
