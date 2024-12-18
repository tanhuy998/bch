package getUserAuthorityDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"app/valueObject"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetUsertAuthorityService struct {
		UserRepo repositoryAPI.ICRUDMongoRepository[model.User] //repository.IUser
	}
)

// func (this *GetUsertAuthorityService) Serve(
// 	tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
// ) (*valueObject.AuthData, error) {
// 	fmt.Println(userUUID)
// 	return this.query(
// 		bson.D{
// 			{"uuid", userUUID},
// 		},
// 		tenantUUID,
// 		ctx,
// 	)
// }

func (this *GetUsertAuthorityService) Serve(
	tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
) (*valueObject.AuthData, error) {

	data, err := repository.AggregateOne[valueObject.AuthData](
		this.UserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"uuid", userUUID},
						{"tenantUUID", tenantUUID},
					},
				},
			},
			// bson.D{
			// 	{
			// 		"$lookup", bson.D{
			// 			{"from", "userSessions"},
			// 			{"localField", "uuid"},
			// 			{"foreignField", "userUUID"},
			// 			{"as", "userSessions"},
			// 			{
			// 				"$pipeline", mongo.Pipeline{
			// 					{
			// 						{
			// 							"$match", bson.D{
			// 								{"uuid", userUUID},
			// 								{"tenantUUID", tenantUUID},
			// 							},
			// 						},
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
			// bson.D{
			// 	{
			// 		"$match", bson.D{
			// 			{
			// 				"userSessions", bson.D{
			// 					{"$ne", bson.A{}},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUsers"},
						{"localField", "uuid"},
						{"foreignField", "userUUID"},
						{"as", "participatedCommandGroups"},
						{"pipeline",
							mongo.Pipeline{
								bson.D{
									{
										"$match", bson.D{
											{"tenantUUID", tenantUUID},
										},
									},
								},
								bson.D{
									{"$lookup",
										bson.D{
											{"from", "commandGroupUserRoles"},
											{"localField", "uuid"},
											{"foreignField", "commandGroupUserUUID"},
											{"as", "roles"},
											{"pipeline",
												mongo.Pipeline{
													bson.D{
														{
															"$match", bson.D{
																{"tenantUUID", tenantUUID},
															},
														},
													},
													bson.D{
														{"$lookup",
															bson.D{
																{"from", "roles"},
																{"localField", "roleUUID"},
																{"foreignField", "uuid"},
																{"as", "details"},
																{"pipeline",
																	mongo.Pipeline{
																		bson.D{{"$project", bson.D{{"name", 1}}}},
																	},
																},
															},
														},
													},
													bson.D{{"$unwind", "$details"}},
													//bson.D{{"$replaceWith", "$details"}},
												},
											},
										},
									},
								},
								bson.D{
									{
										"$set", bson.D{
											{"roles", "$roles.details.name"},
										},
									},
								},
								bson.D{
									{"$project",
										bson.D{
											{"commandGroupUUID", 1},
											{"roles", 1},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "tenantAgents"},
						{"localField", "uuid"},
						{"foreignField", "userUUID"},
						{"as", "tenantAgentData"},
						{
							"pipeline", mongo.Pipeline{
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
				{"$set",
					bson.D{
						{"isTenantAgent",
							bson.D{
								{"$ne",
									bson.A{
										"$tenantAgentData", bson.A{},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"uuid", 1},
						{"name", 1},
						{"username", 1},
						{"tenantUUID", 1},
						{"participatedCommandGroups", 1},
						{"isTenantAgent", 1},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if data == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	if !data.IsTenantAgent() && *data.TenantUUID != tenantUUID {

		return nil, errors.Join(
			common.ERR_FORBIDEN, fmt.Errorf("could no switch to the tenant that the user haven't belonged to"),
		)
	}

	data.Init()

	return data, nil
}
