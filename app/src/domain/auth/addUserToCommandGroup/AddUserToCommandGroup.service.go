package addUserToCommandGroup

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	authServicePort "app/src/port/auth"
	"app/src/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ERR_INVALID_GROUP         = errors.New("invalid group.")
	ERR_INVALID_USER          = errors.New("invalid user.")
	ERR_USER_ALREADY_IN_GROUP = errors.New("user already in group")
)

type (
	AddUserToCommandGroupService struct {
		GetSingleUserService         authServicePort.IGetSingleUser
		GetSingleCommandGroupService authServicePort.IGetSingleCommandGroup
		CheckUserInCommandGroup      authServicePort.ICheckUserInCommandGroup
		CommandGroupUserRepo         repository.ICommandGroupUser
	}
)

func (this *AddUserToCommandGroupService) Get() authServicePort.IGetSingleCommandGroup {

	return this.GetSingleCommandGroupService
}

func (this *AddUserToCommandGroupService) Serve(groupUUID_str string, userUUID_str string) error {

	group, err := this.GetSingleCommandGroupService.Serve(groupUUID_str)

	if err != nil {

		return err
	}

	if group == nil {

		return ERR_INVALID_GROUP
	}

	user, err := this.GetSingleUserService.Serve(userUUID_str, context.TODO())

	if err != nil {

		return err
	}

	if user == nil {

		return ERR_INVALID_USER
	}

	res, err := this.CheckUserInCommandGroup.Detail(*group.UUID, *user.UUID)

	if err != nil {

		return err
	}

	if res != nil {

		return ERR_USER_ALREADY_IN_GROUP
	}

	dataModel := &model.CommandGroupUser{
		UUID:             libCommon.PointerPrimitive(uuid.New()),
		CommandGroupUUID: group.UUID,
		UserUUID:         user.UUID,
	}

	this.CommandGroupUserRepo.Create(dataModel, context.TODO())

	return nil
}
