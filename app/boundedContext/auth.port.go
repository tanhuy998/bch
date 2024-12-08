package boundedContext

import (
	libConfig "app/internal/lib/config"
	"app/model"
	authServicePort "app/port/auth"
	paginateServicePort "app/port/paginate"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"

	removeDBUserSessionDomain "app/domain/auth/RemoveDBUserSession"
	addUserToCommandGroupDomain "app/domain/auth/addUserToCommandGroup"
	checkAuthorityDomain "app/domain/auth/checkAuthority"
	checkCommandGroupUserRolesDomain "app/domain/auth/checkCommandGroupUserRoles"
	checkUserInCommandGroupDomain "app/domain/auth/checkUserInCommandGroup"
	createCommandGroupDomain "app/domain/auth/createCommandGroup"
	createUserDomain "app/domain/auth/createUser"
	getAllRoleDomain "app/domain/auth/getAllRoles"
	"app/domain/auth/getAssignmentGroupUnAssignedCommandGroupUsersDomain"
	getCommandGroupUsersDomain "app/domain/auth/getCommandGroupUsers"
	"app/domain/auth/getSingleCommandGroupDomain"
	getSingleUserDomain "app/domain/auth/getSingleUser"
	getTenantAllGroupsDomain "app/domain/auth/getTenantAllGroups"
	getTenantCommandGroupDomain "app/domain/auth/getTenantCommandGroups"
	getTenantUsersDomain "app/domain/auth/getTenantUsers"
	getUserAuthorityDomain "app/domain/auth/getUserAuthority"
	getUserParticipatedCommandGroupsDomain "app/domain/auth/getUserParticipatedCommandGroups"
	"app/domain/auth/reportUserParticipatedCommandGroupsDomain"

	//getUserParticipatedCommandGroupDomain "app/domain/auth/getUserParticipatedGroups"
	grantCommandGroupRoleToUserDomain "app/domain/auth/grandCommandGroupRolesToUser"
	modifyUserDomain "app/domain/auth/modifyUser"

	"github.com/kataras/iris/v12/hero"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	AuthBoundedContext struct {
		AddUserToCommandGroup            authServicePort.IAddUserToCommandGroup
		CheckCommandGroupUserRole        authServicePort.ICheckCommandGroupUserRole
		CheckUserInCommandGroup          authServicePort.ICheckUserInCommandGroup
		CreateCommandGroup               authServicePort.ICreateCommandGroup
		CreateUser                       authServicePort.ICreateUser
		GetAllRoles                      authServicePort.IGetAllRoles
		GetCommandGroupUsers             authServicePort.IGetCommandGroupUsers
		GetSingleCommandGroup            authServicePort.IGetSingleCommandGroup
		GetSingleUser                    authServicePort.IGetSingleUser
		GrantCommandGroupRolesToUser     authServicePort.IGrantCommandGroupRolesToUser
		ModifyUser                       authServicePort.IModifyUser
		GetUserParticipatedCommandGroups authServicePort.IGetUserParticipatedCommandGroups
	}
)

func RegisterAuthBoundedContext(container *hero.Container) {

	libConfig.BindDependency[authServicePort.IRemoveDBUserSession, removeDBUserSessionDomain.RemoveDBUserSessionService](container, nil)
	libConfig.BindDependency[authServicePort.ICheckUserInCommandGroup, checkUserInCommandGroupDomain.CheckUserInCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICheckCommandGroupUserRole, checkCommandGroupUserRolesDomain.CheckCommandGroupUserRoleService](container, nil)

	libConfig.BindDependency[
		paginateServicePort.IPaginate[model.CommandGroup, primitive.ObjectID],
		getTenantCommandGroupDomain.GetTenantCommandGroupService,
	](container, nil)

	libConfig.BindDependency[authServicePort.IGetAssignmentGroupUnAssignedCommandGroupUsers, getAssignmentGroupUnAssignedCommandGroupUsersDomain.GetAssignmentGroupUnAssignedCommandGroupUserService](container, nil)
	libConfig.BindDependency[authServicePort.IGetUserAuthorityServicePort, getUserAuthorityDomain.GetUsertAuthorityService](container, nil)
	libConfig.BindDependency[authServicePort.IGetAllRoles, getAllRoleDomain.GetAllRolesService](container, nil)
	libConfig.BindDependency[authServicePort.IGetCommandGroupUsers, getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)
	//libConfig.BindDependency[authServicePort.IGetParticipatedCommandGroups, getUserParticipatedCommandGroupDomain.GetParticipatedCommandGroupsService](container, nil)
	//libConfig.BindDependency[]()
	libConfig.BindDependency[authServicePort.IGetTenantUsers[model.User], getTenantUsersDomain.GetTenantUsersService](container, nil)
	libConfig.BindDependency[authServicePort.IGetTenantAllGroups, getTenantAllGroupsDomain.GetTenantAllGroupService](container, nil)

	libConfig.BindDependency[authServicePort.IReportParticipatedCommandGroups, reportUserParticipatedCommandGroupsDomain.ReportParticipatedCommandGroupsService](container, nil)
	libConfig.BindDependency[authServicePort.IGetUserParticipatedCommandGroups, getUserParticipatedCommandGroupsDomain.GetUserParticipatedCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleCommandGroup, getSingleCommandGroupDomain.GetSingleCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleUser, getSingleUserDomain.GetSingleUserService](container, nil)
	libConfig.BindDependency[authServicePort.IGetCommandGroupUsers, getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)

	libConfig.BindDependency[authServicePort.IAddUserToCommandGroup, addUserToCommandGroupDomain.AddUserToCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateCommandGroup, createCommandGroupDomain.CreateCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateUser, createUserDomain.CreateUserService](container, nil)

	libConfig.BindDependency[authServicePort.IGrantCommandGroupRolesToUser, grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserService](container, nil)

	libConfig.BindDependency[authServicePort.IModifyUser, modifyUserDomain.ModifyUserService](container, nil)
	libConfig.BindDependency[authServicePort.ICheckAuthority, checkAuthorityDomain.CheckAuthorityService](container, nil)

	registerDomainSpecificUtils(container)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]],
		getTenantUsersDomain.GetTenantUserUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter],
		createUserDomain.CreateUserUsecase,
	](container, nil)
	// libConfig.BindDependency[
	// 	usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse],
	// 	getGroupUser
	// ]()
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse],
		getCommandGroupUsersDomain.GetCommandGroupUsersUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse],
		modifyUserDomain.ModifyUserUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse],
		getAllRoleDomain.GetAllRolesUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse],
		grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse],
		createCommandGroupDomain.CreateCommandGroupUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse],
		addUserToCommandGroupDomain.AddUserToCommandGroupUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups],
		getUserParticipatedCommandGroupsDomain.GetUserParticipatedCommandGroupsUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.ReportParticipatedGroups, responsePresenter.ReportParticipatedGroups],
		reportUserParticipatedCommandGroupsDomain.ReportParticipatedCommandGroupsUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetTenantAllGroups, responsePresenter.GetTenantAllGroups],
		getTenantAllGroupsDomain.GetTenantAllGroupUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers, responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers],
		getAssignmentGroupUnAssignedCommandGroupUsersDomain.GetAssignmentGroupUnAssignedCommandGroupUsersUseCase,
	](container, nil)

	libConfig.BindDependency[
		usecasePort.IMiddlewareUseCase, checkAuthorityDomain.CheckAuthorityUseCase,
	](container, nil)

	container.Register(new(AuthBoundedContext)).Explicitly().EnableStructDependents()
}
