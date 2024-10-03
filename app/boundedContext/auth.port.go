package boundedContext

import (
	libConfig "app/internal/lib/config"
	authServicePort "app/port/auth"

	addUserToCommandGroupDomain "app/domain/auth/addUserToCommandGroup"
	checkCommandGroupUserRolesDomain "app/domain/auth/checkCommandGroupUserRoles"
	checkUserInCommandGroupDomain "app/domain/auth/checkUserInCommandGroup"
	createCommandGroupDomain "app/domain/auth/createCommandGroup"
	createUserDomain "app/domain/auth/createUser"
	getAllRoleDomain "app/domain/auth/getAllRoles"
	getCommandGroupUsersDomain "app/domain/auth/getCommandGroupUsers"
	getParticipatedCommandGroup "app/domain/auth/getParticipatedCommandGroups"
	"app/domain/auth/getSingleCommandGroupDomain"
	getSingleUserDomain "app/domain/auth/getSingleUser"
	grantCommandGroupRoleToUserDomain "app/domain/auth/grandCommandGroupRolesToUser"
	loginDomain "app/domain/auth/login"
	modifyUserDomain "app/domain/auth/modifyUser"
	refreshLoginDomain "app/domain/auth/refreshLogin"

	"github.com/kataras/iris/v12/hero"
)

type (
	AuthBoundedContext struct {
		authServicePort.IAddUserToCommandGroup
		authServicePort.ICheckCommandGroupUserRole
		authServicePort.ICheckUserInCommandGroup
		authServicePort.ICreateCommandGroup
		authServicePort.ICreateUser
		authServicePort.IGetAllRoles
		authServicePort.IGetCommandGroupUsers
		authServicePort.IGetParticipatedCommandGroups
		authServicePort.IGetSingleCommandGroup
		authServicePort.IGetSingleUser
		authServicePort.IGrantCommandGroupRolesToUser
		authServicePort.ILogIn
		authServicePort.IModifyUser
		authServicePort.IRefreshLogin
	}
)

func RegisterAuthBoundedContext(container *hero.Container) {

	libConfig.BindDependency[authServicePort.ICheckUserInCommandGroup, checkUserInCommandGroupDomain.CheckUserInCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICheckCommandGroupUserRole, checkCommandGroupUserRolesDomain.CheckCommandGroupUserRoleService](container, nil)

	libConfig.BindDependency[authServicePort.IGetAllRoles, getAllRoleDomain.GetAllRolesService](container, nil)
	libConfig.BindDependency[authServicePort.IGetCommandGroupUsers, getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)
	libConfig.BindDependency[authServicePort.IGetParticipatedCommandGroups, getParticipatedCommandGroup.GetParticipatedCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleCommandGroup, getSingleCommandGroupDomain.GetSingleCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleUser, getSingleUserDomain.GetSingleUserService](container, nil)

	libConfig.BindDependency[authServicePort.IAddUserToCommandGroup, addUserToCommandGroupDomain.AddUserToCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateCommandGroup, createCommandGroupDomain.CreateCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateUser, createUserDomain.CreateUserService](container, nil)

	libConfig.BindDependency[authServicePort.IGrantCommandGroupRolesToUser, grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserService](container, nil)
	libConfig.BindDependency[authServicePort.ILogIn, loginDomain.LogInService](container, nil)
	libConfig.BindDependency[authServicePort.IModifyUser, modifyUserDomain.ModifyUserService](container, nil)
	libConfig.BindDependency[authServicePort.IRefreshLogin, refreshLoginDomain.RefreshLoginService](container, nil)

	container.Register(new(AuthBoundedContext)).Explicitly().EnableStructDependents()
}
