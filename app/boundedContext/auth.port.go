package boundedContext

import (
	libConfig "app/internal/lib/config"
	accessTokenServicePort "app/port/accessToken"
	accessTokenClientPort "app/port/accessTokenClient"
	authServicePort "app/port/auth"
	authSignatureTokenPort "app/port/authSignatureToken"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenClientPort "app/port/refreshTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	accessTokenClientService "app/service/accessTokenClient"
	"app/service/accessTokenService"
	"app/service/authSignatureToken"
	refreshTokenService "app/service/refreshToken"
	refreshTokenClientService "app/service/refreshTokenClient"

	addUserToCommandGroupDomain "app/domain/auth/addUserToCommandGroup"
	checkCommandGroupUserRolesDomain "app/domain/auth/checkCommandGroupUserRoles"
	checkUserInCommandGroupDomain "app/domain/auth/checkUserInCommandGroup"
	createCommandGroupDomain "app/domain/auth/createCommandGroup"
	createUserDomain "app/domain/auth/createUser"
	getAllRoleDomain "app/domain/auth/getAllRoles"
	getCommandGroupUsersDomain "app/domain/auth/getCommandGroupUsers"
	"app/domain/auth/getSingleCommandGroupDomain"
	getSingleUserDomain "app/domain/auth/getSingleUser"
	getUserAuthorityDomain "app/domain/auth/getUserAuthority"
	getUserParticipatedCommandGroupDomain "app/domain/auth/getUserParticipatedGroups"
	grantCommandGroupRoleToUserDomain "app/domain/auth/grandCommandGroupRolesToUser"
	loginDomain "app/domain/auth/login"
	modifyUserDomain "app/domain/auth/modifyUser"
	refreshLoginDomain "app/domain/auth/refreshLogin"

	"github.com/kataras/iris/v12/hero"
)

type (
	AuthBoundedContext struct {
		AddUserToCommandGroup        authServicePort.IAddUserToCommandGroup
		CheckCommandGroupUserRole    authServicePort.ICheckCommandGroupUserRole
		CheckUserInCommandGroup      authServicePort.ICheckUserInCommandGroup
		CreateCommandGroup           authServicePort.ICreateCommandGroup
		CreateUser                   authServicePort.ICreateUser
		GetAllRoles                  authServicePort.IGetAllRoles
		GetCommandGroupUsers         authServicePort.IGetCommandGroupUsers
		GetParticipatedCommandGroups authServicePort.IGetParticipatedCommandGroups
		GetSingleCommandGroup        authServicePort.IGetSingleCommandGroup
		GetSingleUser                authServicePort.IGetSingleUser
		GrantCommandGroupRolesToUser authServicePort.IGrantCommandGroupRolesToUser
		LogIn                        authServicePort.ILogIn
		ModifyUser                   authServicePort.IModifyUser
		RefreshLogin                 authServicePort.IRefreshLogin
	}
)

func registerDomainSpecificUtils(container *hero.Container) {

	libConfig.BindDependency[accessTokenServicePort.IAccessTokenManipulator, accessTokenService.JWTAccessTokenManipulatorService](container, nil)
	libConfig.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	refreshTokenService := new(refreshTokenService.RefreshTokenManipulatorService)
	libConfig.BindDependency[refreshTokenServicePort.IRefreshTokenManipulator](container, refreshTokenService)

	libConfig.BindDependency[refreshTokenClientPort.IRefreshTokenClient, refreshTokenClientService.RefreshTokenClientService](container, nil)
	libConfig.BindDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)
}

func RegisterAuthBoundedContext(container *hero.Container) {

	libConfig.BindDependency[authServicePort.ICheckUserInCommandGroup, checkUserInCommandGroupDomain.CheckUserInCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICheckCommandGroupUserRole, checkCommandGroupUserRolesDomain.CheckCommandGroupUserRoleService](container, nil)

	libConfig.BindDependency[authServicePort.IGetUserAuthorityServicePort, getUserAuthorityDomain.GetUsertAuthorityService](container, nil)
	libConfig.BindDependency[authServicePort.IGetAllRoles, getAllRoleDomain.GetAllRolesService](container, nil)
	libConfig.BindDependency[authServicePort.IGetCommandGroupUsers, getCommandGroupUsersDomain.GetCommandGroupUsersService](container, nil)
	libConfig.BindDependency[authServicePort.IGetParticipatedCommandGroups, getUserParticipatedCommandGroupDomain.GetParticipatedCommandGroupsService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleCommandGroup, getSingleCommandGroupDomain.GetSingleCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.IGetSingleUser, getSingleUserDomain.GetSingleUserService](container, nil)

	libConfig.BindDependency[authServicePort.IAddUserToCommandGroup, addUserToCommandGroupDomain.AddUserToCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateCommandGroup, createCommandGroupDomain.CreateCommandGroupService](container, nil)
	libConfig.BindDependency[authServicePort.ICreateUser, createUserDomain.CreateUserService](container, nil)

	libConfig.BindDependency[authServicePort.IGrantCommandGroupRolesToUser, grantCommandGroupRoleToUserDomain.GrantCommandGroupRolesToUserService](container, nil)
	libConfig.BindDependency[authServicePort.ILogIn, loginDomain.LogInService](container, nil)
	libConfig.BindDependency[authServicePort.IModifyUser, modifyUserDomain.ModifyUserService](container, nil)
	libConfig.BindDependency[authServicePort.IRefreshLogin, refreshLoginDomain.RefreshLoginService](container, nil)

	registerDomainSpecificUtils(container)

	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.LoginRequest, responsePresenter.LoginResponse],
		loginDomain.LogInUseCase,
	](container, nil)
	libConfig.BindDependency[
		usecasePort.IUseCase[requestPresenter.RefreshLoginRequest, responsePresenter.RefreshLoginResponse],
		refreshLoginDomain.RefreshLoginUseCase,
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
		usecasePort.IUseCase[requestPresenter.GetParticipatedGroups, responsePresenter.GetParticipatedGroups],
		getUserParticipatedCommandGroupDomain.GetParticipatedCommandGroupsUseCase,
	](container, nil)

	container.Register(new(AuthBoundedContext)).Explicitly().EnableStructDependents()
}
