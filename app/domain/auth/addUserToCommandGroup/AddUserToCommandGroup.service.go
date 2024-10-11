package addUserToCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
	"context"
	"errors"
	"fmt"

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

func (this *AddUserToCommandGroupService) Serve(tenantUUID uuid.UUID, dataModel *model.CommandGroupUser, ctx context.Context) error {

	switch group, err := this.GetSingleCommandGroupService.Serve(*dataModel.CommandGroupUUID, ctx); {
	case err != nil:
		return err
	case group == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("command group not found"))
	case *group.TenantUUID != tenantUUID:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("command group not in tenant"))
	}

	switch user, err := this.GetSingleUserService.Serve(*dataModel.UserUUID, ctx); {
	case err != nil:
		return err
	case user == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case *user.TenantUUID != tenantUUID:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("user not in tenant"))
	}

	switch res, err := this.CheckUserInCommandGroup.Detail(*dataModel.CommandGroupUUID, *dataModel.UserUUID, ctx); {
	case err != nil:
		return err
	case res != nil && *res.TenantUUID == tenantUUID:
		return errors.Join(common.ERR_CONFLICT, fmt.Errorf("user already in group"))
	case res != nil:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("wrong tenant"))
	}

	dataModel.UUID = libCommon.PointerPrimitive(uuid.New())
	dataModel.TenantUUID = libCommon.PointerPrimitive(tenantUUID)

	err := this.CommandGroupUserRepo.Create(dataModel, ctx)

	if err != nil {

		return errors.Join(common.ERR_INTERNAL, err)
	}

	return nil
}
