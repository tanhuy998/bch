package getTenantUsersDomain

import (
	libCommon "app/internal/lib/common"
	libMongo "app/internal/lib/mongo"
	"app/model"
	"app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	RepoPaginateFunc[T libMongo.IBsonDocument] func(
		collection *mongo.Collection, dataModel T, size uint64, ctx context.Context, filters ...primitive.E,
	) ([]T, error)
)

type (
	GetTenantUsersService struct {
		UserRepo repository.IUser
	}
)

func (this *GetTenantUsersService) Serve(
	tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
) ([]model.User, error) {

	if tenantUUID == uuid.Nil {

		return nil, fmt.Errorf("invalid tenant uuid, nil value given")
	}
	fmt.Println(1)
	if cursor == nil {

		return this.UserRepo.FindOffset(
			bson.D{
				{"tenantUUID", tenantUUID},
			},
			page,
			size,
			nil,
			ctx,
		)
	}
	fmt.Println(2)
	dataModel := model.User{}
	dataModel.ObjectID = cursor

	fn := libCommon.Ternary[RepoPaginateFunc[model.User]](
		isPrev,
		repository.FindPrevious,
		repository.FindNext,
	)

	return fn(
		this.UserRepo.GetCollection(),
		dataModel,
		size,
		ctx,
		bson.E{"tenantUUID", tenantUUID},
	)
}
