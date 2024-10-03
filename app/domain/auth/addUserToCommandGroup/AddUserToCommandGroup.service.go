package addUserToCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
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

func (this *AddUserToCommandGroupService) Serve(groupUUID_str uuid.UUID, userUUID_str uuid.UUID, ctx context.Context) error {

	group, err := this.GetSingleCommandGroupService.Serve(groupUUID_str, ctx)

	if err != nil {

		return err
	}

	if group == nil {

		return ERR_INVALID_GROUP
	}

	user, err := this.GetSingleUserService.Serve(userUUID_str, ctx)

	if err != nil {

		return err
	}

	if user == nil {

		return errors.Join(common.ERR_NOT_FOUND, errors.New("user not found"))
	}

	res, err := this.CheckUserInCommandGroup.Detail(*group.UUID, *user.UUID, ctx)

	if err != nil {

		return err
	}

	if res != nil {

		return errors.New("user already in group") // ERR_USER_ALREADY_IN_GROUP
	}

	dataModel := &model.CommandGroupUser{
		UUID:             libCommon.PointerPrimitive(uuid.New()),
		CommandGroupUUID: group.UUID,
		UserUUID:         user.UUID,
	}

	err = this.CommandGroupUserRepo.Create(dataModel, ctx)

	if err != nil {

		return errors.Join(common.ERR_INTERNAL, err)
	}

	return nil
}
