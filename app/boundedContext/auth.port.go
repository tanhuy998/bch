package boundedContext

import (
	"app/domain"
	irisIoc "app/internal/lib/iris/ioc"
	"app/model"
	authServicePort "app/port/auth"
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
)

type (
	AuthBoundedContext struct {
		AddUserToCommandGroupService            authServicePort.IAddUserToCommandGroup
		CheckCommandGroupUserRoleService        authServicePort.ICheckCommandGroupUserRole
		CheckUserInCommandGroupService          authServicePort.ICheckUserInCommandGroup
		CreateCommandGroupService               authServicePort.ICreateCommandGroup
		CreateUserService                       authServicePort.ICreateUser
		GetAllRolesService                      authServicePort.IGetAllRoles
		GetCommandGroupUsersService             authServicePort.IGetCommandGroupUsers[domain.PaginateCursorType]
		GetSingleCommandGroupService            authServicePort.IGetSingleCommandGroup
		GetSingleUserService                    authServicePort.IGetSingleUser
		GrantCommandGroupRolesToUserService     authServicePort.IGrantCommandGroupRolesToUser
		ModifyUserService                       authServicePort.IModifyUser
		GetUserParticipatedCommandGroupsService authServicePort.IGetUserParticipatedCommandGroups
	}
)

func RegisterAuthBoundedContext(container *hero.Container) {

	irisIoc.BindDependency[authServicePort.IRemoveDBUserSession, removeDBUserSessionDomain.RemoveDBUserSessionService](container, nil)
	irisIoc.BindDependency[authServicePort.ICheckUserInCommandGroup, checkUserInCommandGroupDomain.CheckUserInCommandGroupService](container, nil)
	irisIoc.BindDependency[authServicePort.ICheckCommandGroupUserRole, checkCommandGroupUserRolesDomain.CheckCommandGroupUserRoleService](container, nil)

	irisIoc.BindDependency[
		//paginateServicePort.IPaginateService[model.CommandGroup, primitive.ObjectID],
		authServicePort.IGetTenantCommandGroups[model.CommandGroup],
		getTenantCommandGroupDomain.GetTenantCommandGroupService,
	](container, nil)

	irisIoc.BindDependency[authServicePort.IGetAssignmentGroupUnAssignedCommandGroupUsers, getAssignmentGroupUnAssignedCommandGroupUsersDomain.GetAssignmentGroupUnAssignedCommandGroupUserService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetUserAuthorityServicePort, getUserAuthorityDomain.GetUsertAuthorityService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetAllRoles, getAllRoleDomain.GetAllRolesService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetCommandGroupUsers[domain.PaginateCursorType], getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)
	//libConfig.BindDependency[authServicePort.IGetParticipatedCommandGroups, getUserParticipatedCommandGroupDomain.GetParticipatedCommandGroupsService](container, nil)
	//libConfig.BindDependency[]()
	irisIoc.BindDependency[authServicePort.IGetTenantUsers[model.User], getTenantUsersDomain.GetTenantUsersService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetTenantAllGroups, getTenantAllGroupsDomain.GetTenantAllGroupService](container, nil)

	irisIoc.BindDependency[authServicePort.IReportParticipatedCommandGroups, reportUserParticipatedCommandGroupsDomain.ReportParticipatedCommandGroupsService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetUserParticipatedCommandGroups, getUserParticipatedCommandGroupsDomain.GetUserParticipatedCommandGroupService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetSingleCommandGroup, getSingleCommandGroupDomain.GetSingleCommandGroupService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetSingleUser, getSingleUserDomain.GetSingleUserService](container, nil)
	irisIoc.BindDependency[authServicePort.IGetCommandGroupUsers[domain.PaginateCursorType], getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)

	irisIoc.BindDependency[authServicePort.IAddUserToCommandGroup, addUserToCommandGroupDomain.AddUserToCommandGroupService](container, nil)
	irisIoc.BindDependency[authServicePort.ICreateCommandGroup, createCommandGroupDomain.CreateCommandGroupService](container, nil)
	irisIoc.BindDependency[authServicePort.ICreateUser, createUserDomain.CreateUserService](container, nil)

	irisIoc.BindDependency[authServicePort.IGrantCommandGroupRolesToUser, grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserService](container, nil)

	irisIoc.BindDependency[authServicePort.IModifyUser, modifyUserDomain.ModifyUserService](container, nil)
	irisIoc.BindDependency[authServicePort.ICheckAuthority, checkAuthorityDomain.CheckAuthorityService](container, nil)

	registerDomainSpecificUtils(container)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetTenantCommandGroups, responsePresenter.GetTenantCommandGroups[model.CommandGroup]],
		getTenantCommandGroupDomain.GetTenantCommandGroupsUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]],
		getTenantUsersDomain.GetTenantUserUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter],
		createUserDomain.CreateUserUsecase,
	](container, nil)
	// libConfig.BindDependency[
	// 	usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse],
	// 	getGroupUser
	// ]()
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse],
		getCommandGroupUsersDomain.GetCommandGroupUsersUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.ModifyUserRequest, responsePresenter.ModifyUserResponse],
		modifyUserDomain.ModifyUserUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAllRolesRequest, responsePresenter.GetAllRolesResponse],
		getAllRoleDomain.GetAllRolesUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GrantCommandGroupRolesToUserRequest, responsePresenter.GrantCommandGroupRolesToUserResponse],
		grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.CreateCommandGroupRequest, responsePresenter.CreateCommandGroupResponse],
		createCommandGroupDomain.CreateCommandGroupUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse],
		addUserToCommandGroupDomain.AddUserToCommandGroupUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetUserParticipatedCommandGroups, responsePresenter.GetUserParticipatedCommandGroups],
		getUserParticipatedCommandGroupsDomain.GetUserParticipatedCommandGroupsUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.ReportParticipatedGroups, responsePresenter.ReportParticipatedGroups],
		reportUserParticipatedCommandGroupsDomain.ReportParticipatedCommandGroupsUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetTenantAllGroups, responsePresenter.GetTenantAllGroups],
		getTenantAllGroupsDomain.GetTenantAllGroupUseCase,
	](container, nil)
	irisIoc.BindDependency[
		usecasePort.IUseCase[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers, responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers],
		getAssignmentGroupUnAssignedCommandGroupUsersDomain.GetAssignmentGroupUnAssignedCommandGroupUsersUseCase,
	](container, nil)

	irisIoc.BindDependency[
		usecasePort.IMiddlewareUseCase, checkAuthorityDomain.CheckAuthorityUseCase,
	](container, nil)

	container.Register(new(AuthBoundedContext)).Explicitly().EnableStructDependents()
}
