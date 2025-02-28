package getSingleAssignmentDomain

import (
	"app/internal/common"
	"app/internal/db/query"
	"app/model"
	"app/repository"
	repositoryAPI "app/repository/api"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type (
	GetSingleAssignmentService struct {
		GetSingleAssigmentAggregate
		AssignmentRepo repository.IAssignment
	}
)

func (this *GetSingleAssignmentService) ServeUsingRelations(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context,
) (*model.Assignment, error) {

	return this.GetSingleAssigmentAggregate.
		MergeRelations().
		Filter(
			func(filter query.IFilterExpression) {

				filter.Field("tenantUUID").Equal(tenantUUID)
				filter.Field("uuid").Equal(assignmentUUID)
			},
		).First(ctx)
}

func (this *GetSingleAssignmentService) Serve(
	tenantUUID uuid.UUID, uuid uuid.UUID, ctx context.Context,
) (*model.Assignment, error) {

	// ret, err := repository.AggregateOne[model.Assignment](
	// 	this.AssignmentRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 					{"uuid", uuid},
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
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$lookup", bson.D{
	// 					{"from", "assignmentGroups"},
	// 					{"localField", "uuid"},
	// 					{"foreignField", "assignmentUUID"},
	// 					{"as", "assignmentGroups"},
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
	// 	},
	// 	ctx,
	// )

	ret, err := this.ServeUsingRelations(tenantUUID, uuid, ctx)

	if err != nil {

		return nil, err
	}

	if ret == nil {

		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	}

	// if ret.TenantUUID == nil {

	// 	return nil, libError.NewInternal(fmt.Errorf("wrong data"))
	// }

	// if *ret.TenantUUID != tenantUUID {

	// 	return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	// }

	return ret, nil
}

func (this *GetSingleAssignmentService) Search(
	model *model.Assignment, ctx context.Context,
) (*model.Assignment, error) {

	// ret, err := this.AssignmentRepo.Find(
	// 	bson.D{
	// 		{"title", model.Title},
	// 		{"tenantUUID", model.TenantUUID},
	// 	},
	// 	ctx,
	// )

	ret, err := this.AssignmentRepo.Filter(
		func(filter repositoryAPI.IFilterGenerator) {
			filter.Field("tenantUUID").Equal(model.TenantUUID)
			filter.Field("title").Equal(model.Title)
		},
	).FindOne(ctx)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return ret, nil
}
