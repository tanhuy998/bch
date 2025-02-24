package getAssignmentGroupsDomain

import (
	"app/domain"
	"app/internal/common"
	"app/internal/db/query"
	"app/model"
	paginateServicePort "app/port/paginate"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type (
	GetAssignmentGroupsService struct {
		AssignmentGroupRepo repository.IAssignmentGroup
		//AssignmentGroupAggregate aggregate.AggegateRoot[model.AssignmentGroup, model.AssignmentGroup]
		GetAssignmentGroupAggegate
	}
)

// func (this *GetAssignmentGroupsService) TestAggregate(
// 	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator domain.IPaginator, ctx context.Context,
// ) ([]model.AssignmentGroup, error) {

// 	return this.AssignmentGroupAggregate.Query().
// 		Join("users", func(queryBuilder query.IJoinField) {
// 			queryBuilder.On("createdBy", "uuid").As("createdUser")
// 		}).
// 		Join("commandGroups", func(queryBuilder query.IJoinField) {
// 			queryBuilder.On("commandGroupUUID", "uuid").As("commandGroup")
// 		}).
// 		ExcludeFields(
// 			"user.password",
// 			"user.secret",
// 		).
// 		Paginate(ctx, func() paginateServicePort.IPaginator[interface{}] {

// 			return (paginator).(paginateServicePort.IPaginator[interface{}])
// 		})
// }

func (this *GetAssignmentGroupsService) Serve(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, paginator domain.IPaginator, ctx context.Context,
) ([]model.AssignmentGroup, error) {

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

	return this.GetAssignentGroups().
		Filter(
			func(filter query.IFilterExpression) {

				filter.Field("tenantUUID").Equal(tenantUUID)
			},
		).
		ExcludeFields(
			"user.password",
			"user.secret",
		).
		Paginate(
			ctx,
			func() paginateServicePort.IPaginator[interface{}] {

				return (paginator).(paginateServicePort.IPaginator[interface{}])
			},
		)
}
