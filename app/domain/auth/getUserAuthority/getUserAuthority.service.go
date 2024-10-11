package getUserAuthorityDomain

import (
	"app/internal/common"
	"app/repository"
	"app/valueObject"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetUsertAuthorityService struct {
		UserRepo repository.IUser
	}
)

func (this *GetUsertAuthorityService) Serve(
	tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
) (*valueObject.AuthData, error) {

	return this.query(
		bson.D{
			{"uuid", userUUID},
		},
		tenantUUID,
		ctx,
	)
}

func (this *GetUsertAuthorityService) query(
	criterias bson.D, tenantUUID uuid.UUID, ctx context.Context,
) (*valueObject.AuthData, error) {

	data, err := repository.AggregateOne[valueObject.AuthData](
		this.UserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", criterias,
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUsers"},
						{"localField", "uuid"},
						{"foreignField", "userUUID"},
						{"as", "participatedCommandGroups"},
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
				{"$match",
					bson.D{
						{"$or",
							bson.A{
								bson.D{{"isTenantAgent", true}},
								bson.D{{"tenantUUID", "$tenantAgentData.0.uuid"}},
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

	return data, nil
}
