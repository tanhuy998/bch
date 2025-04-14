package getAssignmentGroupsDomain

import (
	"app/internal/common"
	"app/internal/db/query"
	"app/model"
	assignmentServicePort "app/port/assignment"
	paginateServicePort "app/port/paginate"

	"app/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetAssignmentGroupsService struct {
		AssignmentGroupRepo repository.IAssignmentGroup
		GetAssignmentGroupAggegate
	}
)

func (this *GetAssignmentGroupsService) Serve(
	//tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator domain.IPaginator, ctx context.Context,
	ctx assignmentServicePort.IGetAssignmentInputContext[primitive.ObjectID],
) ([]model.AssignmentGroup, error) {

	var (
		tenantUUID     = ctx.GetTenantUUID()
		assignmentUUID = ctx.GetAssignmentUUID()
		paginator      = ctx.GetPaginator()
	)

	switch {
	case tenantUUID == uuid.Nil:
		return nil, common.ERR_UNAUTHORIZED
	case assignmentUUID == uuid.Nil:
		return nil, errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("invalid assignment uuid"))
	}

	// return repository.Aggregate[model.AssignmentGroup](
	// 	this.AssignmentGroupRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "users"},
	// 					{"localField", "createdBy"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "createdUser"},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "commandGroups"},
	// 					{"localField", "commandGroupUUID"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "commandGroup"},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$set",
	// 				bson.D{
	// 					{"commandGroup",
	// 						bson.D{
	// 							{"$arrayElemAt",
	// 								bson.A{
	// 									"$commandGroup",
	// 									0,
	// 								},
	// 							},
	// 						},
	// 					},
	// 					{"createdUser",
	// 						bson.D{
	// 							{"$arrayElemAt",
	// 								bson.A{
	// 									"$createdUser",
	// 									0,
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$project",
	// 				bson.D{
	// 					{"user",
	// 						bson.D{
	// 							{"password", 0},
	// 							{"secret", 0},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	ctx,
	// )

	// return this.GetAssignentGroups().
	// 	Filter(
	// 		func(filter query.IFilterExpression) {

	// 			filter.Field("tenantUUID").Equal(tenantUUID)
	// 		},
	// 	).
	// 	ExcludeFields(
	// 		"user.password",
	// 		"user.secret",
	// 	).
	// 	Paginate(
	// 		ctx,
	// 		func() paginateServicePort.IPaginator[interface{}] {

	// 			return (paginator).(paginateServicePort.IPaginator[interface{}])
	// 		},
	// 	)

	res, err := this.ByDomain(ctx).Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.
				Filter(
					func(filter query.IFilterExpression) {

						filter.Field("tenantUUID").Equal(tenantUUID)
						filter.Field("assignmentUUID").Equal(assignmentUUID)
					},
				).
				ExcludeFields(
					"createdUser.password",
					"createdUser.username",
				)
		},
	).Paginate(
		(paginator).(paginateServicePort.IPaginator[interface{}]),
		ctx,
	)

	fmt.Println(err)

	return res, err
}
