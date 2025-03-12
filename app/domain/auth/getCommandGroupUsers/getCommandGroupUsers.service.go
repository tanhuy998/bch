package getCommandGroupUsersDomain

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

var (
	ERR_INVALID_INPUT_GROUP = errors.New("GetGroupUsers: invalid group")
	ERR_NO_GROUP            = errors.New("GetGroupUser: command group not found")
)

type (
	IGetCommandGroupUsers interface {
		Serve(groupUUID string) ([]*model.User, error)
	}

	GetCommandGroupUsersService struct {
		GetCommandGroupUserAggregate
		CommandGroupUserRepo repository.ICommandGroupUser
		CommandGroupRepo     repository.ICommandGroup
	}
)

func (this *GetCommandGroupUsersService) ServeUsingRelation(
	tenantUUID uuid.UUID, groupUUID uuid.UUID, paginator domain.Paginator, ctx context.Context,
) ([]model.CommandGroupUser, error) {

	// return this.GetCommandGroupUserAggregate.
	// 	MergeRelations().
	// 	Filter(
	// 		func(filter query.IFilterExpression) {
	// 			filter.Field("tenantUUID").Equal(tenantUUID)
	// 			filter.Field("commandGroupUUID").Equal(groupUUID)
	// 			filter.Field("user").Not().Equal(nil)
	// 		},
	// 	).
	// 	Paginate(
	// 		ctx,
	// 		func() paginateServicePort.IPaginator[interface{}] {

	// 			return paginateServicePort.IPaginator[interface{}](paginator)
	// 		},
	// 	)

	return this.GetCommandGroupUserAggregate.MergeRelations().Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.
				Filter(
					func(filter query.IFilterExpression) {
						filter.Field("tenantUUID").Equal(tenantUUID)
						filter.Field("commandGroupUUID").Equal(groupUUID)
						filter.Field("user").Not().Equal(nil)
					},
				)
		},
	).Paginate(
		(paginator).(paginateServicePort.IPaginator[interface{}]),
		ctx,
	)
}

func (this *GetCommandGroupUsersService) Serve(
	tenantUUID uuid.UUID, groupUUID uuid.UUID, paginator domain.Paginator, ctx context.Context,
) ([]model.CommandGroupUser, error) {

	switch existingGroup, err := this.CommandGroupRepo.FindOneByUUID(groupUUID, ctx); {
	case err != nil:
		return nil, err
	case existingGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("group not found"))
	case *existingGroup.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("group not in tenant"))
	}

	// ret, err := repository.Aggregate[model.CommandGroupUser](
	// 	this.CommandGroupUserRepo.GetCollection(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"tenantUUID", tenantUUID},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "users"},
	// 					{"localField", "userUUID"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "users"},
	// 					{
	// 						"pipeline", mongo.Pipeline{
	// 							bson.D{
	// 								{
	// 									"$project", bson.D{
	// 										{"name", 1},
	// 										{"username", 1},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		bson.D{{"$unwind", bson.D{{"path", "$users"}}}},
	// 		bson.D{
	// 			{
	// 				"$set", bson.D{
	// 					{"user", "$users"},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{
	// 				"$project", bson.D{
	// 					{"users", 0},
	// 				},
	// 			},
	// 		},
	// 		//bson.D{{"$replaceRoot", bson.D{{"newRoot", "$$ROOT.users"}}}},
	// 	},
	// 	ctx,
	// )

	// if err != nil {

	// 	return nil, err
	// }

	// return ret, nil

	return this.ServeUsingRelation(tenantUUID, groupUUID, paginator, ctx)
}

// func (this *GetCommandGroupUsersService) SearchAndRetrieveByModel(dataModel *model.CommandGroup, ctx context.Context) ([]*model.User, error) {

// 	groupUUID := dataModel.UUID
// 	group, err := this.CommandGroupRepo.FindOneByUUID(*groupUUID, context.TODO())

// 	if err != nil {

// 		return nil, err
// 	}

// 	if group == nil {

// 		return nil, ERR_NO_GROUP
// 	}

// 	ret, err := repository.Aggregate[model.User](
// 		this.CommandGroupUserRepo.GetCollection(),
// 		mongo.Pipeline{
// 			bson.D{
// 				{
// 					"$match", dataModel,
// 				},
// 			},
// 			bson.D{
// 				{"$lookup",
// 					bson.D{
// 						{"from", "users"},
// 						{"localField", "userUUID"},
// 						{"foreignField", "uuid"},
// 						{"as", "users"},
// 					},
// 				},
// 			},
// 			bson.D{{"$unwind", bson.D{{"path", "$users"}}}},
// 			bson.D{{"$replaceRoot", bson.D{{"newRoot", "$$ROOT.users"}}}},
// 		},
// 		context.TODO(),
// 	)

// 	if err != nil {

// 		return nil, err
// 	}

// 	return ret, nil
// }
