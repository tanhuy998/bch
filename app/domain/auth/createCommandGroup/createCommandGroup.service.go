package createCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
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
		CommandGroupRepo             repository.ICommandGroup
		GetSingleCommandGroupService authServicePort.IGetSingleCommandGroup
	}
)

func (this *CreateCommandGroupService) Serve(groupName string, ctx context.Context) error {

	groupExist, err := this.GetSingleCommandGroupService.CheckCommandGroupExistence(groupName, ctx)

	if err != nil {

		return err
	}

	if groupExist {

		return ERR_COMMAND_EXIST
	}

	model := &model.CommandGroup{
		UUID: libCommon.PointerPrimitive(uuid.New()),
		Name: groupName,
	}

	err = this.CommandGroupRepo.Create(model, ctx)

	if err != nil {

		return err
	}

	return nil
}

func (this *CreateCommandGroupService) CreateByModel(model *model.CommandGroup, ctx context.Context) (*model.CommandGroup, error) {

	groupExist, err := this.CommandGroupRepo.Find(
		bson.D{
			{"name", model.Name},
			{"tenantUUID", model.TenantUUID},
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

	err = this.CommandGroupRepo.Create(model, ctx)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return model, nil
}
