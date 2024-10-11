package createCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ERR_COMMAND_EXIST = errors.New("command group already exists.")
)

type (
	CreateCommandGroupService struct {
		CommandGroupRepo repository.ICommandGroup
	}
)

func (this *CreateCommandGroupService) Serve(tenantUUID uuid.UUID, groupName string, ctx context.Context) (*model.CommandGroup, error) {

	model := &model.CommandGroup{
		Name: groupName,
	}

	return this.CreateByModel(tenantUUID, model, ctx)
}

func (this *CreateCommandGroupService) CreateByModel(
	tenantUUID uuid.UUID, model *model.CommandGroup, ctx context.Context,
) (*model.CommandGroup, error) {

	if tenantUUID == uuid.Nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	groupExist, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", model.Name},
			{"tenantUUID", tenantUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if groupExist != nil {

		return nil, ERR_COMMAND_EXIST
	}

	model.UUID = libCommon.PointerPrimitive(uuid.New())
	model.TenantUUID = libCommon.PointerPrimitive(tenantUUID)

	err = this.CommandGroupRepo.Create(model, ctx)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return model, nil
}
