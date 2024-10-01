package createCommandGroupDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	authServicePort "app/src/port/auth"
	"app/src/repository"
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

func (this *CreateCommandGroupService) Serve(groupName string) error {

	groupExist, err := this.GetSingleCommandGroupService.CheckCommandGroupExistence(groupName)

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

	return this.CommandGroupRepo.Create(model, context.TODO())
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

		return nil, err
	}

	return model, nil
}
