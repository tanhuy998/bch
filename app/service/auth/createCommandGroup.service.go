package authService

import (
	"app/domain/model"
	"app/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_COMMAND_EXIST = errors.New("command group already exists.")
)

type (
	ICreateCommandGroup interface {
		Serve(groupName string) error
	}

	CreateCommandGroupService struct {
		CommandGroupRepo             repository.ICommandGroup
		GetSingleCommandGroupService IGetSingleCommandGroup
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
		UUID: uuid.New(),
		Name: groupName,
	}

	return this.CommandGroupRepo.Create(model, context.TODO())
}
