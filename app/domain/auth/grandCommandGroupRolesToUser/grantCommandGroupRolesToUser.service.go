package grantCommandGroupRoleToUserDomain

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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ERR_INVALID_GROUP_USER          = errors.New("GrantCommandGroupUserRole: invalid group user.")
	ERR_EMPTY_ROLE_LIST             = errors.New("GrantCommandGroupUserRole: empty role list.")
	ERR_INVALID_VALUES_IN_ROLE_LIST = errors.New("GrantCommandGroupUserRole: invalid values in role list.")
	ERR_ROLES_EXIST                 = errors.New("GrantCommandGroupUserRole: roles already granted.")
)

type (
	IGrantCommandGroupRolesToUser interface {
		Serve(groupUUID string, userUUID string, roles []string) error
	}

	GrantCommandGroupRolesToUserService struct {
		RoleRepo                         repository.IRole
		CheckCommandGroupUserRoleService authServicePort.ICheckCommandGroupUserRole
		CommandGroupUserRepo             repository.ICommandGroupUser
		CommandGroupUserRoleRepo         repository.ICommandGroupUserRole
		GetSingleCommandGroupService     authServicePort.IGetSingleCommandGroup
		CheckUserInCommandGroup          authServicePort.ICheckUserInCommandGroup
	}
)

func (this *GrantCommandGroupRolesToUserService) Serve(
	groupUUID_str string,
	userUUID_str string,
	roles []string,
) error {

	if len(roles) == 0 {

		return ERR_EMPTY_ROLE_LIST
	}

	groupUUID, err := uuid.Parse(groupUUID_str)

	if err != nil {

		return errors.Join(common.ERR_BAD_REQUEST, errors.New("invalid group"))
	}

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		return errors.Join(common.ERR_BAD_REQUEST, errors.New("invalid user"))
	}

	commandGroupUser, err := this.CheckUserInCommandGroup.Detail(groupUUID, userUUID)

	if err != nil {

		return err
	}

	if commandGroupUser == nil {

		return ERR_INVALID_GROUP_USER
	}

	roleUUIDList, err := this._checkValidRoles(roles)

	if err != nil {

		return err
	}

	if len(roleUUIDList) == 0 {

		return ERR_INVALID_VALUES_IN_ROLE_LIST
	}

	unGrantedRoles, err := this.CheckCommandGroupUserRoleService.Compare(groupUUID, userUUID, roleUUIDList)

	if err != nil {

		return err
	}

	if len(unGrantedRoles) == 0 {

		return ERR_ROLES_EXIST
	}

	var commandGroupUserRoleList []*model.CommandGroupUserRole = make([]*model.CommandGroupUserRole, len(unGrantedRoles))

	for i, v := range unGrantedRoles {

		commandGroupUserRoleList[i] = &model.CommandGroupUserRole{
			UUID:                 libCommon.PointerPrimitive(uuid.New()),
			RoleUUID:             &v,
			CommandGroupUserUUID: commandGroupUser.UUID,
		}
	}

	err = this.CommandGroupUserRoleRepo.CreateMany(commandGroupUserRoleList, context.TODO())

	if err != nil {

		return err
	}

	return nil
}

func (this *GrantCommandGroupRolesToUserService) _checkValidRoles(roleUUIDs []string) ([]uuid.UUID, error) {

	var (
		conditions []bson.D    = make([]primitive.D, len(roleUUIDs))
		ret        []uuid.UUID = make([]uuid.UUID, len(roleUUIDs))
	)

	for i, v := range roleUUIDs {

		roleUUID, err := uuid.Parse(v)

		if err != nil {

			return nil, ERR_INVALID_VALUES_IN_ROLE_LIST
		}

		ret[i] = roleUUID
		conditions[i] = bson.D{{"uuid", roleUUID}}
	}

	res, err := this.RoleRepo.FindMany(
		bson.D{
			{"$or", conditions},
		},
		context.TODO(),
	)

	if err != nil || res == nil || len(res) != len(roleUUIDs) {

		return nil, ERR_INVALID_VALUES_IN_ROLE_LIST
	}

	return ret, nil
}
