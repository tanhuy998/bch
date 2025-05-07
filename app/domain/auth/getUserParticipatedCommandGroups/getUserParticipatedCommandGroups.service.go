package getUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	"app/internal/db/driver/mongoDriver/mongoStorage"
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
	"context"
	"errors"
	"fmt"
)

type (
	GetUserParticipatedCommandGroupService struct {
		UserRepo repository.IUser
		Stu      mongoStorage.IMongoDBStorageUnit[model.CommandGroupUser]
		//CommandGroupUserRepo repository.ICommandGroupUser
		domain_aggregate
	}
)

func (this *GetUserParticipatedCommandGroupService) Serve(
	//tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
	input authServicePort.GetUserParticipatedCommandGroupInput,
) ([]model.CommandGroup, error) {

	var (
		userUUID   = input.GetRequestedUserUUID()
		tenantUUID = input.GetTenantUUID()
		ctx        = input.GetContext()
	)

	switch existingUser, err := this.UserRepo.FindOneByUUID(userUUID, ctx); {
	case err != nil:
		return nil, err
	case existingUser == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case *existingUser.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("user not in tenant"))
	}

	// _, err := repository.Aggregate[model.CommandGroup](
	// 	this.Stu.GetStorageUnit(),
	// 	mongo.Pipeline{
	// 		bson.D{
	// 			{
	// 				"$match", bson.D{
	// 					{"userUUID", userUUID},
	// 					{"tenantUUID", tenantUUID},
	// 				},
	// 			},
	// 		},
	// 		bson.D{
	// 			{"$lookup",
	// 				bson.D{
	// 					{"from", "commandGroups"},
	// 					{"localField", "commandGroupUUID"},
	// 					{"foreignField", "uuid"},
	// 					{"as", "commandGroups"},
	// 				},
	// 			},
	// 		},
	// 		bson.D{{"$unwind", "$commandGroups"}},
	// 		bson.D{{"$replaceWith", "$commandGroups"}},
	// 	},
	// 	ctx,
	// )

	ret, err := this.MergeRelations(tenantUUID, userUUID).
		Read(nil).
		Paginate(input.GetGeneralPaginator(), ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *GetUserParticipatedCommandGroupService) SearchAndRetrieveByModel(
	searchModel *model.User, ctx context.Context,
) ([]model.CommandGroup, error) {

	return nil, nil
}
