package grantCommandGroupRoleToUserDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	"app/model"
	authServicePort "app/port/auth"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ERR_INVALID_GROUP_USER          = errors.New("GrantCommandGroupUserRole: invalid group user.")
	ERR_EMPTY_ROLE_LIST             = errors.New("GrantCommandGroupUserRole: empty role list.")
	ERR_INVALID_VALUES_IN_ROLE_LIST = errors.New("GrantCommandGroupUserRole: invalid values in role list.")
	ERR_ROLES_EXIST                 = errors.New("GrantCommandGroupUserRole: roles already granted.")
)

type (
	GrantCommandGroupRolesToUserService struct {
		RoleRepo                         repository.IRole
		UserRepo                         repository.IUser
		CheckCommandGroupUserRoleService authServicePort.ICheckCommandGroupUserRole
		CommandGroupUserRepo             repository.ICommandGroupUser
		CommandGroupUserRoleRepo         repository.ICommandGroupUserRole
		GetSingleCommandGroupService     authServicePort.IGetSingleCommandGroup
		CheckUserInCommandGroup          authServicePort.ICheckUserInCommandGroup
		MongoClient                      *mongo.Client
	}
)

func (this *GrantCommandGroupRolesToUserService) Serve(
	tenantUUID uuid.UUID,
	groupUUID uuid.UUID,
	userUUID uuid.UUID,
	roles []uuid.UUID,
	createdBy uuid.UUID,
	ctx context.Context,
) error {

	if len(roles) == 0 {

		return errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("no roles provided"))
	}

	switch existingUser, err := this.UserRepo.FindOneByUUID(userUUID, ctx); {
	case err != nil:
		return err
	case existingUser == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case *existingUser.TenantUUID != tenantUUID:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("user not in tenant"))
	}

	existingCommandGroupUser, err := this.CheckUserInCommandGroup.Detail(groupUUID, userUUID, ctx)

	if err != nil && !errors.Is(err, common.ERR_NOT_FOUND) {

		return err
	}

	if err = this.checkValidRoles(roles, ctx); err != nil {

		return err
	}

	var (
		unGrantedRoles      []uuid.UUID
		newCommandGroupUser *model.CommandGroupUser
	)

	if existingCommandGroupUser == nil {
		// if error was not added to group, prepare new data model
		// and create later
		unGrantedRoles = roles

		newCommandGroupUser = &model.CommandGroupUser{
			UUID:             libCommon.PointerPrimitive(uuid.New()),
			UserUUID:         &userUUID,
			CommandGroupUUID: &groupUUID,
		}

	} else {

		unGrantedRoles, err = this.CheckCommandGroupUserRoleService.Compare(groupUUID, userUUID, roles, ctx)

		if err != nil {

			return err
		}
	}

	if len(unGrantedRoles) == 0 {

		return errors.Join(common.ERR_CONFLICT, fmt.Errorf("given roles already grant to user"))
	}

	var commandGroupUserRoleList []*model.CommandGroupUserRole = make([]*model.CommandGroupUserRole, len(unGrantedRoles))

	for i, v := range unGrantedRoles {

		obj := &model.CommandGroupUserRole{
			UUID:                 libCommon.PointerPrimitive(uuid.New()),
			RoleUUID:             &v,
			CommandGroupUserUUID: libCommon.Ternary(existingCommandGroupUser != nil, existingCommandGroupUser.UUID, newCommandGroupUser.UUID),
			TenantUUID:           &tenantUUID,
		}

		if createdBy != uuid.Nil {

			obj.CreatedBy = &createdBy
		}

		commandGroupUserRoleList[i] = obj
	}

	session, err := this.MongoClient.StartSession()

	if err != nil {

		return libError.NewInternal(err)
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(
		ctx, func(sessionCtx mongo.SessionContext) (interface{}, error) {

			if existingCommandGroupUser == nil {

				err := this.CommandGroupUserRepo.Create(newCommandGroupUser, sessionCtx)

				if err != nil {

					return nil, err
				}
			}

			err = this.CommandGroupUserRoleRepo.CreateMany(commandGroupUserRoleList, sessionCtx)

			if err != nil {

				return nil, err
			}

			return nil, nil
		},
	)

	if err != nil {

		return err
	}

	return nil
}

func (this *GrantCommandGroupRolesToUserService) checkValidRoles(roleUUIDs []uuid.UUID, ctx context.Context) error {

	var (
		conditions []bson.D    = make([]primitive.D, len(roleUUIDs))
		ret        []uuid.UUID = make([]uuid.UUID, len(roleUUIDs))
	)

	for i, v := range roleUUIDs {

		ret[i] = v
		conditions[i] = bson.D{{"uuid", v}}
	}

	res, err := this.RoleRepo.FindMany(
		bson.D{
			{"$or", conditions},
		},
		ctx,
	)

	if err != nil || res == nil || len(res) != len(roleUUIDs) {

		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("invalid roles provided"))
	}

	return nil
}
