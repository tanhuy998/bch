package navigateTenantDomain

import (
	"app/internal/db/query"
	"app/model"
	authGenServicePort "app/port/authGenService"
	"app/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	NavigateTenantService struct {
		domain_aggregate
		TenantRepo      repository.ITenant
		UserSessionRepo repository.IUserSession
	}
)

func (this *NavigateTenantService) Serve(
	// userUUID uuid.UUID, ctx context.Context
	input authGenServicePort.NavigateTenantInput,
) ([]model.Tenant, error) {

	// data, err := repository.Aggregate[model.Tenant](
	// 	this.TenantRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "tenantAgents"},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "tenantUUID"},
	// 					{"as", "tenantAgent"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{"userUUID", userUUID},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$unwind",
	// 				bson.D{
	// 					{"path", "$tenantAgent"},
	// 					{"preserveNullAndEmptyArrays", true},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "users"},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "tenantUUID"},
	// 					{"as", "user"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$match", bson.D{
	// 										{
	// 											"$and", bson.A{
	// 												bson.D{
	// 													{
	// 														"uuid", bson.D{
	// 															{"$eq", userUUID},
	// 														},
	// 													},
	// 												},
	// 												bson.D{
	// 													{
	// 														"uuid", bson.D{
	// 															{"$ne", "$tenantAgent.userUUID"},
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
	// 			{"$unwind",
	// 				bson.D{
	// 					{"path", "$user"},
	// 					{"preserveNullAndEmptyArrays", true},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{
	// 						"isTenantAgent", bson.D{
	// 							{
	// 								"$cond", bson.D{
	// 									{
	// 										"if", bson.D{
	// 											{
	// 												"$ne", bson.A{"$tenantAgent", nil},
	// 											},
	// 										},
	// 									},
	// 									{"then", true},
	// 									{"else", false},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$project", bson.D{
	// 					{"description", 0},
	// 					{"tenantAgent", 0},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	data, err := this.MergeRelations(
		input.GetRequestedUserUUID(),
	).Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.Transform(
				func(it query.IDataTransformer) {

					it.Set("isTenantAgent").Value(
						bson.D{
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
					)
				},
			).ExcludeFields(
				"description", "tenantAgent",
			)
		},
	).All(input.GetContext())

	if err != nil {

		return nil, err
	}

	fmt.Println(len(data))

	return data, err
}
