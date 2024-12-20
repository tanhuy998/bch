package config

import (
	"app/boundedContext"
	"app/infrastructure/http/common"
	"app/internal/bootstrap"
	"app/internal/db"
	"app/internal/generalToken"
	libConfig "app/internal/lib/config"
	accessLogServicePort "app/port/accessLog"
	actionResultServicePort "app/port/actionResult"
	cacheListServicePort "app/port/cacheList"
	dbQueryTracerPort "app/port/dbQueryTracer"
	generalTokenServicePort "app/port/generalToken"
	generalTokenClientServicePort "app/port/generalTokenClient"
	generalTokenIDServicePort "app/port/generalTokenID"
	"log"

	jwtTokenServicePort "app/port/jwtTokenService"
	passwordServicePort "app/port/passwordService"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	"app/port/responsePresetPort"
	uniqueIDServicePort "app/port/uniqueID"
	"app/repository"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"
	cacheListService "app/service/cacheList"
	generalTokenClientService "app/service/generalTokenClient"
	generalTokenIDService "app/service/generalTokenID"
	"app/service/generalTokenService"
	irisAccessLoggerService "app/service/irisAccessLogger"
	jwtTokenService "app/service/jwtToken"
	"app/service/mongoDBTracerService"
	passwordService "app/service/password"
	refreshTokenIDService "app/service/refreshTokenID"
	"app/service/responsePresetService"
	uniqueIDService "app/service/uniqueID"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

const (
	VALIDATOR    = "validator"
	AUTH         = "auth_service"
	DBMS_CLIENT  = "dbms_client"
	DB           = "db_instancce"
	REQUEST_BODY = "request_body"
)

func InitializeDatabase(app router.Party) {

	var container *hero.Container = app.ConfigureContainer().EnableStructDependents().Container

	fmt.Println("Initialize DBMS client...")
	client := db.GetClient()

	db := db.GetDB()

	// container.Register(func(ctx iris.Context) *mongo.Client {

	// 	ctx.Values().Set(DBMS_CLIENT, client)

	// 	return client
	// })
	// container.Register(func(ctx iris.Context) *mongo.Database {

	// 	ctx.Values().Set(DB, db)

	// 	return db
	// })

	container.Register(log.Default()).Explicitly()
	container.Register(db).Explicitly()
	container.Register(client).Explicitly()

	fmt.Println("DBMS client initialized.")

	/*
		access logger must be initialized before repositories in order to trace db query
	*/
	libConfig.BindDependency[
		accessLogServicePort.IAccessLogger,
		irisAccessLoggerService.IrisAccessLoggerService,
	](container, nil)

	libConfig.BindDependency[
		dbQueryTracerPort.IDBQueryTracer, mongoDBTracerService.DBQueryTracerService,
	](container, nil)

	fmt.Println("Initialize Repositories...")
	libConfig.BindDependency[repository.ITenant](
		container, new(repository.TenantRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ITenantAgent](
		container, new(repository.TenantAgentRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IUser](
		container, new(repository.UserRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICommandGroup](
		container, new(repository.CommandGroupRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICommandGroupUser](
		container, new(repository.CommandGroupUserRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICommandGroupUserRole](
		container, new(repository.CommandGroupUserRoleRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IRole](
		container, new(repository.RoleRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICampaignRepository](
		container, new(repository.CampaignRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICandidateRepository](
		container, new(repository.CandidateRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICandidateSigningCommit](
		container, new(repository.CandidateSingingCommitRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.ICandidateSigningInfo](
		container, new(repository.CandidateSigningInfoRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IAssignment](
		container, new(repository.AssignmentRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IAssignmentGroup](
		container, new(repository.AssignmentGroupRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IAssignmentGroupMember](
		container, new(repository.AssignmentGroupMemberRepository).Init(db),
	).EnableStructDependents()
	libConfig.BindDependency[repository.IUserSession](
		container, new(repository.UserSessionRepository).Init(db),
	).EnableStructDependents()
	fmt.Println("Repositories Initialized.")
}

// func RegisterAdapters(container *hero.Container) {

// 	fmt.Println("Wiring dependencies adapters...")

// 	libConfig.BindDependency[adminServiceAdapter.ICheckCandidateExistence, adminService.CheckCandidateExistenceService](container, nil)
// 	libConfig.BindDependency[adminServiceAdapter.IGetSingleCandidate, adminService.AdminGetSingleCandidateByUUIDService](container, nil)
// 	libConfig.BindDependency[signingServiceAdapter.ICountSignedCandidates, signingService.CountSignedCandidateService](container, nil)
// 	libConfig.BindDependency[signingServiceAdapter.IGetCampaignSignedCandidates, signingService.GetCampaignSignedCandidates](container, nil)
// 	libConfig.BindDependency[signingServiceAdapter.IGetCampaignUnSignedCandidates, signingService.GetCampaignUnSignedCandidatesService](container, nil)

// 	libConfig.BindDependency[authServiceAdapter.IGetSingleUserService, authService.GetSingleUser](container, nil)
// 	libConfig.BindDependency[authServiceAdapter.ICreateUserService, authService.CreateUserService](container, nil)
// 	//libConfig.BindDependency[tenantAgentServiceAdapter.IGetSingleTenantAgentServiceAdapter, tenantService.GetSingleTenantAgentService](container, nil)

// 	fmt.Println("Wiring dependencies adapters successfully.")
// }

// func RegisterTenantDependencies(container *hero.Container) {

// 	libConfig.BindDependency[tenantServicePort.ICreateTenantAgent, tenantService.CreateTenantAgentService](container, nil)
// 	libConfig.BindDependency[tenantServicePort.ICreateTenant, tenantService.CreateTenantService](container, nil)

// 	libConfig.BindDependency[usecase.ICreateTenant, usecase.CreateTenantUseCase](container, nil)
// }

// func RegisterAuthEndpointServiceDependencies(container *hero.Container) {

// 	db := db.GetDB()

// 	authService.Initialize(db)

// 	libConfig.BindDependency[authService.IAuthService, authService.AuthenticationService](container, nil)

// 	libConfig.BindDependency[authService.IGetAllRoles, authService.GetAllRolesService](container, nil)

// 	libConfig.BindDependency[authService.IGetSingleUser, authService.GetSingleUser](container, nil)
// 	libConfig.BindDependency[authService.IGetSingleCommandGroup, authService.GetSingleCommandGroupService](container, nil)
// 	libConfig.BindDependency[authService.ICheckUserInCommandGroup, authService.CheckUserInCommandGroupService](container, nil)
// 	libConfig.BindDependency[authService.ICheckCommandGroupUserRole, authService.CheckCommandGroupUserRoleService](container, nil)
// 	libConfig.BindDependency[authService.IGetCommandGroupUsers, authService.GetCommandGroupUsersService](container, nil)

// 	libConfig.BindDependency[authService.ICreateUser, authService.CreateUserService](container, nil)
// 	libConfig.BindDependency[authService.ICreateCommandGroup, authService.CreateCommandGroupService](container, nil)
// 	libConfig.BindDependency[authService.IAddUserToCommandGroup, authService.AddUserToCommandGroupService](container, nil)
// 	libConfig.BindDependency[authService.IGetParticipatedCommandGroups, authService.GetParticipatedCommandGroupsService](container, nil)
// 	libConfig.BindDependency[authService.IGrantCommandGroupRolesToUser, authService.GrantCommandGroupRolesToUserService](container, nil)
// 	libConfig.BindDependency[authService.ICheckCommandGroupUserRole, authService.CheckCommandGroupUserRoleService](container, nil)
// 	libConfig.BindDependency[authService.IModifyUser, authService.ModifyUserService](container, nil)
// 	libConfig.BindDependency[authServiceAdapter.ILogIn, authService.LogInService](container, nil)
// 	libConfig.BindDependency[authServiceAdapter.IRefreshLogin, authService.RefreshLoginService](container, nil)

// 	libConfig.BindDependency[usecase.ICreateUser, usecase.CreateUserUsecase](container, nil)

// 	libConfig.BindDependency[usecase.ICreateCommandGroup, usecase.CreateCommandGroupUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IAddUserToCommandGroup, usecase.AddUserToCommandGroupUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IGetParticipatedCommandGroups, usecase.GetParticipatedCommandGroupsUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IGrantCommandGroupRolesToUser, usecase.GrantCommandGroupRolesToUserUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IGetAllRoles, usecase.GetAllRolesUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IGetGroupUsers, usecase.GetGroupUsersUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IModifyUser, usecase.ModifyUserUseCase](container, nil)
// 	libConfig.BindDependency[usecase.ILogIn, usecase.LogInUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IRefreshLogin, usecase.RefreshLoginUseCase](container, nil)
// }

// func RegisterAssignmentEndpointServiceDependencies(container *hero.Container) {

// 	libConfig.BindDependency[assignmentServicePort.IGetSingleAssignnment, assignmentService.GetSingleAssignmentService](container, nil)
// 	libConfig.BindDependency[assignmentServicePort.ICreateAssignment, assignmentService.CreateAssignmentService](container, nil)

// 	libConfig.BindDependency[usecase.ICreateAssignment, usecase.CreateAssignmentUseCase](container, nil)
// 	libConfig.BindDependency[usecase.IGetSingleAssignment, usecase.GetSingleAssignmentUseCase](container, nil)
// }

func RegisterUtilServices(container *hero.Container) {

	libConfig.BindDependency[context.Validator, validator.Validate](container, validator.New())
	libConfig.BindDependency[actionResultServicePort.IActionResult, actionResultService.ResponseResultService](container, nil)
	libConfig.BindDependency[responsePresetPort.IResponsePreset, responsePresetService.ResponsePresetService](container, nil)
	libConfig.BindDependency[passwordServicePort.IPassword, passwordService.PasswordService](container, nil)
	libConfig.BindDependency[common.IMiddlewareErrorHandler, common.ErrorHandler](container, nil)
	// container.Register(new(common.Controller)).Explicitly().EnableStructDependents()
}

func RegisterCaches(container *hero.Container) {

	refreshTokenBlackListCacheClient, err := bootstrap.NewRefreshTokenBlackListClient()

	if err != nil {

		panic("error while inittiating refresh token blacklist cache client: " + err.Error())
	}

	container.Register(refreshTokenBlackListCacheClient)

	generalTokenWhiteListCacheClient, err := bootstrap.NewGeneralTokenWhiteListClient()

	if err != nil {

		panic("error while inittiating general token whitelist cache client: " + err.Error())
	}

	container.Register(generalTokenWhiteListCacheClient)

	libConfig.BindDependency[
		cacheListServicePort.ICacheList[string, bootstrap.RefreshTokenBlackListCacheValue],
		cacheListService.CacheListManipulator[string, bootstrap.RefreshTokenBlackListCacheValue],
	](
		container,
		cacheListService.NewCacheListManipulator[string, bootstrap.RefreshTokenBlackListCacheValue]("refresh_token_black_list"),
	)

	libConfig.BindDependency[
		cacheListServicePort.ICacheList[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue],
		cacheListService.CacheListManipulator[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue],
	](
		container,
		cacheListService.NewCacheListManipulator[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue]("general_token_white_list"),
	)
}

func RegisterAuthDependencies(container *hero.Container) {

	fmt.Println("Initialize Auth service...")
	//BindDependency[authService.IAuthService, authService.AuthenticateService](container, nil)
	// auth := new(authService.AuthenticateService)
	// container.Register(func(ctx iris.Context) authService.IAuthService {

	// 	ctx.Values().Set(AUTH, auth)

	// 	return auth
	// })

	libConfig.BindAndMapDependencyToContext[authService.IAuthService, authService.AuthenticationService](container, nil, AUTH)

	asymmetricJWTService := jwtTokenService.NewECDSAService(
		jwt.SigningMethodES256, *bootstrap.GetJWTAsymmetricEncryptionPrivateKey(), *bootstrap.GetJWTAsymmetricEncryptionPublicKey(),
	)
	libConfig.BindDependency[jwtTokenServicePort.IAsymmetricJWTTokenManipulator](container, asymmetricJWTService)

	symmetricJWTService := jwtTokenService.NewHMACService(
		jwt.SigningMethodHS256, bootstrap.GetJWTSymmetricEncryptionSecret(),
	)
	libConfig.BindDependency[jwtTokenServicePort.ISymmetricJWTTokenManipulator](container, symmetricJWTService)

	uniqueID, err := uniqueIDService.New(15)

	if err != nil {

		panic("error while initiating uniqueID service: " + err.Error())
	}

	libConfig.BindDependency[uniqueIDServicePort.IUniqueIDGenerator](container, uniqueID)
	libConfig.BindDependency[generalTokenIDServicePort.IGeneralTokenIDProvider, generalTokenIDService.GeneralTokenIDProvider](container, nil)

	libConfig.BindDependency[refreshTokenIdServicePort.IRefreshTokenIDProvider, refreshTokenIDService.RefreshTokenIDProviderService](container, nil)

	libConfig.BindDependency[generalTokenServicePort.IGeneralTokenManipulator, generalTokenService.GeneralTokenManipulator](container, nil)
	libConfig.BindDependency[generalTokenClientServicePort.IGeneralTokenClient, generalTokenClientService.GeneralTokenClientService](container, nil)

	//accessTokenSevice := new(accessTokenService.JWTAccessTokenManipulatorService)

	// libConfig.BindDependency[accessTokenServicePort.IAccessTokenManipulator, accessTokenService.JWTAccessTokenManipulatorService](container, nil)
	// libConfig.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	// cacheClient, err := memoryCache.NewClient[string, refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload](
	// 	refreshTokenBlackListServicePort.REFRESH_TOKE_BLACK_LIST_TOPIC,
	// )

	// if err != nil {

	// 	panic("error while inittiating refresh token blacklist cache client: " + err.Error())
	// }

	// libConfig.BindDependency[refreshTokenBlackListServicePort.IRefreshTokenCacheClient](container, cacheClient)
	// libConfig.BindDependency[refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator, refreshTokenBlackListService.RefreshTokenBlackListManipulatorService](container, nil)

	// refreshTokenService := new(refreshTokenService.RefreshTokenManipulatorService)
	// libConfig.BindDependency[refreshTokenServicePort.IRefreshTokenManipulator](container, refreshTokenService)

	// libConfig.BindDependency[refreshTokenClientPort.IRefreshTokenClient, refreshTokenClientService.RefreshTokenClientService](container, nil)
	// libConfig.BindDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)

	// libConfig.BindDependency[accessTokenServicePort.IAccessTokenManipulator, accessTokenService.JWTAccessTokenManipulatorService](container, nil)
	// libConfig.BindDependency[accessTokenClientPort.IAccessTokenClient, accessTokenClientService.BearerAccessTokenClientService](container, nil)

	// //refreshTokenService := new(refreshTokenService.RefreshTokenManipulatorService)
	// libConfig.BindDependency[refreshTokenServicePort.IRefreshTokenManipulator, refreshTokenService.RefreshTokenManipulatorService](container, nil)

	// libConfig.BindDependency[refreshTokenClientPort.IRefreshTokenClient, refreshTokenClientService.RefreshTokenClientService](container, nil)
	// libConfig.BindDependency[authSignatureTokenPort.IAuthSignatureProvider, authSignatureToken.AuthSignatureTokenService](container, nil)

	fmt.Println("Auth service initialized.")
}

func RegisterServices(app router.Party) {

	var container *hero.Container = app.ConfigureContainer().EnableStructDependents().Container

	fmt.Println("Wiring dependencies...")

	RegisterUtilServices(container)
	// RegisterAdapters(container)
	RegisterCaches(container)
	RegisterAuthDependencies(container)
	boundedContext.RegisterAuthBoundedContext(container)
	boundedContext.RegisterTenantBoundedContext(container)
	boundedContext.RegisterAuthGenBoundedContext(container)
	boundedContext.RegisterAuthSignaturesBoundedContext(container)
	boundedContext.RegisterAssignmentBoundedContext(container)
	// RegisterTenantDependencies(container)
	// RegisterAuthEndpointServiceDependencies(container)
	// RegisterAssignmentEndpointServiceDependencies(container)

	// auth := new(authService.AuthenticateService)
	// Dep := container.Register(auth)
	// Dep.DestType = libCommon.InterfaceOf((*authService.IAuthService)(nil))

	//libConfig.BindDependency[port.IActionResult, usecase.ActionResultUseCase](container, nil)

	// libConfig.BindDependency[signingService.ICountSignedCandidate, signingService.CountSignedCandidateService](container, nil)

	/*
		init app validator
	*/

	// /*
	// 	Bind Admin Campaign controller dependent services
	// */
	// libConfig.BindDependency[adminService.IGetCampaign, adminService.AdminGetCampaignService](container, nil)
	// libConfig.BindDependency[adminService.IDeleteCampaign, adminService.AdminDeleteCampaignService](container, nil)
	// libConfig.BindDependency[adminService.ILaunchNewCampaign, adminService.AdminLaunchNewCampaignService](container, nil)
	// libConfig.BindDependency[adminService.IModifyExistingCampaign, adminService.AdminModifyExistingCampaign](container, nil)
	// libConfig.BindDependency[adminService.IGetCampaignList, adminService.AdminGetCampaignListService](container, nil)
	// libConfig.BindDependency[adminService.IGetPendingCampaigns, adminService.AdminGetPendingCampaigns](container, nil)

	// /*
	// 	Bind Admin Candidate controller dependent services
	// */
	// libConfig.BindDependency[adminService.IDeleteCandidate, adminService.AdminDeleteCandidateService](container, nil)
	// libConfig.BindDependency[adminService.IAddNewCandidate, adminService.AdminAddNewCandidateToCampaign](container, nil)
	// libConfig.BindDependency[adminService.IModifyExistingCandidate, adminService.AdminModifyExistingCandidate](container, nil)
	// libConfig.BindDependency[adminService.IGetCampaignCandidateList, adminService.AdminGetCampaignCandidateListService](container, nil)
	// libConfig.BindDependency[adminService.IGetSingleCandidateByUUID, adminService.AdminGetSingleCandidateByUUIDService](container, nil)

	// libConfig.BindDependency[adminService.IGetCampaignSignedCandidates, adminService.GetCampaignSignedCandidates](container, nil)
	// libConfig.BindDependency[adminService.IGetCampaignUnSignedCandidates, adminService.GetCampaignUnSignedCandidatesService](container, nil)

	// libConfig.BindDependency[adminService.ICandidateSigningReport, adminService.CandidateSigningReportService](container, nil)

	// libConfig.BindDependency[candidateService.ICommitCandidateSigningInfo, candidateService.CommitCandidateSigningInfoService](container, nil)
	// libConfig.BindDependency[candidateService.IGetSingleCandidateSigningInfo, candidateService.GetSingleCandidateSigningInfoService](container, nil)
	// libConfig.BindDependency[candidateService.ICheckSigningExistence, candidateService.CheckSigningExistenceService](container, nil)

	// libConfig.BindDependency[candidateService.ICandidateSigningCommitLogger, candidateService.CandidateSigningCommmitLoggerService](container, nil)

	// /*
	// 	Bind Signing Module Services
	// */
	// libConfig.BindDependency[signingService.ICheckCandidateExistence, signingService.CheckCandidateExistenceService](container, nil)
	// libConfig.BindDependency[signingService.ISigningCommitLogger, signingService.SigningCommmitLoggerService](container, nil)
	// libConfig.BindDependency[signingService.ICommitSpecificSigningInfo, signingService.CommitSpecificSigningInfoService](container, nil)
	// //libConfig.BindDependency[signingService.IGetCampaignSignedCandidates, signingService.IGetCampaignSignedCandidates](container, nil)
	// libConfig.BindDependency[signingService.IGetSingleCandidateSigningInfo, signingService.GetSingleCandidateSigningInfoService](container, nil)

	// /*
	// 	Bind Usecase Objects
	// */
	// libConfig.BindDependency[usecase.IGetSingleCampaign, usecase.GetSingleCampaignUseCase](container, nil)
	// libConfig.BindDependency[usecase.ILaunchNewCampaign, usecase.LaunchNewCampaignUseCase](container, nil)
	// libConfig.BindDependency[usecase.IUpdateCampaign, usecase.UpdateCampaignUseCase](container, nil)
	// libConfig.BindDependency[usecase.IDeleteCampaign, usecase.DeleteCampaignUseCase](container, nil)
	// libConfig.BindDependency[usecase.IGetCampaignList, usecase.GetCampaignListUseCase](container, nil)
	// libConfig.BindDependency[usecase.IGetPendingCampaigns, usecase.GetPendingCampaignsUseCase](container, nil)

	// libConfig.BindDependency[usecase.IAddNewCandidate, usecase.AddNewCandidateUseCase](container, nil)
	// libConfig.BindDependency[usecase.IModifyExistingCandidate, usecase.ModifyExistingCandidateUseCase](container, nil)
	// libConfig.BindDependency[usecase.IDeleteCandidate, usecase.DeleteCandidateUseCase](container, nil)

	// libConfig.BindDependency[usecase.IGetCampaignCandidateList, usecase.GetCampaignCandidateListUseCase](container, nil)
	// libConfig.BindDependency[usecase.IGetSingleCandidateByUUID, usecase.GetSingleCandidateByUUIDUseCase](container, nil)

	// libConfig.BindDependency[usecase.ICommitCandidateSigningInfo, usecase.CommitCandidateSigningInfoUseCase](container, nil)
	// libConfig.BindDependency[usecase.IGetSingleCandidateSigningInfo, usecase.GetSingleCandidateSigningInfoUseCase](container, nil)

	// libConfig.BindDependency[usecase.ICheckSigningExistence, usecase.CheckSigningExistenceUseCase](container, nil)
	// libConfig.BindDependency[usecase.IGetCampaignSignedCandidates, usecase.GetCampaignSignedCandidatesUseCase](container, nil)

	// libConfig.BindDependency[usecase.IGetCampaignUnSignedCandidates, usecase.GetCampaignUnSignedCandidatesUseCase](container, nil)

	// libConfig.BindDependency[usecase.ICampaignProgress, usecase.CampaignProgressUseCase](container, nil)

	// libConfig.BindDependency[usecase.ICommitSpecificSigningInfo, usecase.CommitSpecificSigningInfoUseCase](container, nil)

	fmt.Println("Wiring dependencies successully.")
}
