package navigateTenantDomain

import (
	"app/model"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	NavigateTenantService struct {
		TenantRepo repository.ITenant
	}
)

func (this *NavigateTenantService) Serve(userUUID uuid.UUID, ctx context.Context) ([]*model.Tenant, error) {

	data, err := repository.Aggregate[model.Tenant](
		this.TenantRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "tenantAgents"},
						{"localField", "uuid"},
						{"foreignField", "tenantUUID"},
						{"as", "tenantAgent"},
						{
							"pipeline", mongo.Pipeline{
								bson.D{
									{
										"$match", bson.D{
											{"userUUID", userUUID},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$tenantAgent"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "users"},
						{"localField", "uuid"},
						{"foreignField", "tenantUUID"},
						{"as", "user"},
						{
							"pipeline", mongo.Pipeline{
								bson.D{
									{
										"$match", bson.D{
											{
												"$and", bson.A{
													bson.D{
														{
															"uuid", bson.D{
																{"$eq", userUUID},
															},
														},
													},
													bson.D{
														{
															"uuid", bson.D{
																{"$ne", "$tenantAgent.userUUID"},
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
					},
				},
			},
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$user"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			},
			bson.D{
				{
					"$set", bson.D{
						{
							"isTenantAgent", bson.D{
								{
									"$cond", bson.D{
										{
											"if", bson.D{
												{
													"$ne", bson.A{"$tenantAgent", nil},
												},
											},
										},
										{"then", true},
										{"else", false},
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
						{"description", 0},
						{"tenantAgent", 0},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	return data, err
}
