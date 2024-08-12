package config

import (
	adminServiceAdapter "app/adapter/adminService"
	passwordServiceAdapter "app/adapter/passwordService"
	signingServiceAdapter "app/adapter/signingService"
	"app/infrastructure/db"
	libConfig "app/lib/config"
	"app/repository"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	authService "app/service/auth"
	candidateService "app/service/candidate"
	passwordService "app/service/password"
	"app/service/signingService"
	usecase "app/useCase"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	VALIDATOR    = "validator"
	AUTH         = "auth_service"
	DBMS_CLIENT  = "dbms_client"
	DB           = "db_instancce"
	REQUEST_BODY = "request_body"
)

func InitializeDatabase(app router.Party) {

	var container *hero.Container = app.ConfigureContainer().Container

	fmt.Println("Initialize DBMS client...")
	client, err := db.GetClient()

	if err != nil {

		panic(err)
	}

	db := db.GetDB()

	container.Register(func(ctx iris.Context) *mongo.Client {

		ctx.Values().Set(DBMS_CLIENT, client)

		return client
	})
	container.Register(func(ctx iris.Context) *mongo.Database {

		ctx.Values().Set(DB, db)

		return db
	})
	fmt.Println("DBMS client initialized.")

	fmt.Println("Initialize Repositories...")
	libConfig.BindDependency[repository.IUser](
		container, new(repository.UserRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICommandGroup](
		container, new(repository.CommandGroupRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICommandGroupUser](
		container, new(repository.CommandGroupUserRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICommandGroupUserRole](
		container, new(repository.CommandGroupUserRoleRepository).Init(db),
	)
	libConfig.BindDependency[repository.IRole](
		container, new(repository.RoleRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICampaignRepository](
		container, new(repository.CampaignRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICandidateRepository](
		container, new(repository.CandidateRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICandidateSigningCommit](
		container, new(repository.CandidateSingingCommitRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICandidateSigningInfo](
		container, new(repository.CandidateSigningInfoRepository).Init(db),
	)
	fmt.Println("Repositories Initialized.")
}

func RegisterAdapters(container *hero.Container) {

	fmt.Println("Wiring dependencies adapters...")

	libConfig.BindDependency[adminServiceAdapter.ICheckCandidateExistence, adminService.CheckCandidateExistenceService](container, nil)
	libConfig.BindDependency[adminServiceAdapter.IGetSingleCandidate, adminService.AdminGetSingleCandidateByUUIDService](container, nil)
	libConfig.BindDependency[signingServiceAdapter.ICountSignedCandidates, signingService.CountSignedCandidateService](container, nil)
	libConfig.BindDependency[signingServiceAdapter.IGetCampaignSignedCandidates, signingService.GetCampaignSignedCandidates](container, nil)
	libConfig.BindDependency[signingServiceAdapter.IGetCampaignUnSignedCandidates, signingService.GetCampaignUnSignedCandidatesService](container, nil)
	libConfig.BindDependency[passwordServiceAdapter.IPassword, passwordService.PasswordService](container, nil)

	fmt.Println("Wiring dependencies adapters successfully.")
}

func RegisterAuthServices(container *hero.Container) {

	db := db.GetDB()

	authService.Initialize(db)

	libConfig.BindDependency[authService.IAuthService, authService.AuthenticationService](container, nil)

	libConfig.BindDependency[authService.IGetSingleUser, authService.GetSingleUser](container, nil)
	libConfig.BindDependency[authService.ICreateUser, authService.CreateUserService](container, nil)

	libConfig.BindDependency[usecase.ICreateUser, usecase.CreateUserUsecase](container, nil)
}

func RegisterUtilServices(container *hero.Container) {
	libConfig.BindDependency[context.Validator, validator.Validate](container, validator.New())
	libConfig.BindDependency[actionResultService.IActionResult, actionResultService.ResponseResultService](container, nil)
}

func RegisterServices(app router.Party) {

	var container *hero.Container = app.ConfigureContainer().EnableStructDependents().Container

	RegisterUtilServices(container)
	RegisterAdapters(container)
	RegisterAuthServices(container)

	// auth := new(authService.AuthenticateService)
	// Dep := container.Register(auth)
	// Dep.DestType = libCommon.InterfaceOf((*authService.IAuthService)(nil))
	fmt.Println("Initialize Auth service...")
	//BindDependency[authService.IAuthService, authService.AuthenticateService](container, nil)
	// auth := new(authService.AuthenticateService)
	// container.Register(func(ctx iris.Context) authService.IAuthService {

	// 	ctx.Values().Set(AUTH, auth)

	// 	return auth
	// })

	libConfig.BindAndMapDependencyToContext[authService.IAuthService, authService.AuthenticationService](container, nil, AUTH)

	fmt.Println("Auth service initialized.")

	//libConfig.BindDependency[port.IActionResult, usecase.ActionResultUseCase](container, nil)

	fmt.Println("Wiring dependencies...")

	libConfig.BindDependency[signingService.ICountSignedCandidate, signingService.CountSignedCandidateService](container, nil)

	/*
		init app validator
	*/

	/*
		Bind Admin Campaign controller dependent services
	*/
	libConfig.BindDependency[adminService.IGetCampaign, adminService.AdminGetCampaignService](container, nil)
	libConfig.BindDependency[adminService.IDeleteCampaign, adminService.AdminDeleteCampaignService](container, nil)
	libConfig.BindDependency[adminService.ILaunchNewCampaign, adminService.AdminLaunchNewCampaignService](container, nil)
	libConfig.BindDependency[adminService.IModifyExistingCampaign, adminService.AdminModifyExistingCampaign](container, nil)
	libConfig.BindDependency[adminService.IGetCampaignList, adminService.AdminGetCampaignListService](container, nil)
	libConfig.BindDependency[adminService.IGetPendingCampaigns, adminService.AdminGetPendingCampaigns](container, nil)

	/*
		Bind Admin Candidate controller dependent services
	*/
	libConfig.BindDependency[adminService.IDeleteCandidate, adminService.AdminDeleteCandidateService](container, nil)
	libConfig.BindDependency[adminService.IAddNewCandidate, adminService.AdminAddNewCandidateToCampaign](container, nil)
	libConfig.BindDependency[adminService.IModifyExistingCandidate, adminService.AdminModifyExistingCandidate](container, nil)
	libConfig.BindDependency[adminService.IGetCampaignCandidateList, adminService.AdminGetCampaignCandidateListService](container, nil)
	libConfig.BindDependency[adminService.IGetSingleCandidateByUUID, adminService.AdminGetSingleCandidateByUUIDService](container, nil)

	libConfig.BindDependency[adminService.IGetCampaignSignedCandidates, adminService.GetCampaignSignedCandidates](container, nil)
	libConfig.BindDependency[adminService.IGetCampaignUnSignedCandidates, adminService.GetCampaignUnSignedCandidatesService](container, nil)

	libConfig.BindDependency[adminService.ICandidateSigningReport, adminService.CandidateSigningReportService](container, nil)

	libConfig.BindDependency[candidateService.ICommitCandidateSigningInfo, candidateService.CommitCandidateSigningInfoService](container, nil)
	libConfig.BindDependency[candidateService.IGetSingleCandidateSigningInfo, candidateService.GetSingleCandidateSigningInfoService](container, nil)
	libConfig.BindDependency[candidateService.ICheckSigningExistence, candidateService.CheckSigningExistenceService](container, nil)

	libConfig.BindDependency[candidateService.ICandidateSigningCommitLogger, candidateService.CandidateSigningCommmitLoggerService](container, nil)

	/*
		Bind Signing Module Services
	*/
	libConfig.BindDependency[signingService.ICheckCandidateExistence, signingService.CheckCandidateExistenceService](container, nil)
	libConfig.BindDependency[signingService.ISigningCommitLogger, signingService.SigningCommmitLoggerService](container, nil)
	libConfig.BindDependency[signingService.ICommitSpecificSigningInfo, signingService.CommitSpecificSigningInfoService](container, nil)
	//libConfig.BindDependency[signingService.IGetCampaignSignedCandidates, signingService.IGetCampaignSignedCandidates](container, nil)
	libConfig.BindDependency[signingService.IGetSingleCandidateSigningInfo, signingService.GetSingleCandidateSigningInfoService](container, nil)

	/*
		Bind Usecase Objects
	*/
	libConfig.BindDependency[usecase.IGetSingleCampaign, usecase.GetSingleCampaignUseCase](container, nil)
	libConfig.BindDependency[usecase.ILaunchNewCampaign, usecase.LaunchNewCampaignUseCase](container, nil)
	libConfig.BindDependency[usecase.IUpdateCampaign, usecase.UpdateCampaignUseCase](container, nil)
	libConfig.BindDependency[usecase.IDeleteCampaign, usecase.DeleteCampaignUseCase](container, nil)
	libConfig.BindDependency[usecase.IGetCampaignList, usecase.GetCampaignListUseCase](container, nil)
	libConfig.BindDependency[usecase.IGetPendingCampaigns, usecase.GetPendingCampaignsUseCase](container, nil)

	libConfig.BindDependency[usecase.IAddNewCandidate, usecase.AddNewCandidateUseCase](container, nil)
	libConfig.BindDependency[usecase.IModifyExistingCandidate, usecase.ModifyExistingCandidateUseCase](container, nil)
	libConfig.BindDependency[usecase.IDeleteCandidate, usecase.DeleteCandidateUseCase](container, nil)

	libConfig.BindDependency[usecase.IGetCampaignCandidateList, usecase.GetCampaignCandidateListUseCase](container, nil)
	libConfig.BindDependency[usecase.IGetSingleCandidateByUUID, usecase.GetSingleCandidateByUUIDUseCase](container, nil)

	libConfig.BindDependency[usecase.ICommitCandidateSigningInfo, usecase.CommitCandidateSigningInfoUseCase](container, nil)
	libConfig.BindDependency[usecase.IGetSingleCandidateSigningInfo, usecase.GetSingleCandidateSigningInfoUseCase](container, nil)

	libConfig.BindDependency[usecase.ICheckSigningExistence, usecase.CheckSigningExistenceUseCase](container, nil)
	libConfig.BindDependency[usecase.IGetCampaignSignedCandidates, usecase.GetCampaignSignedCandidatesUseCase](container, nil)

	libConfig.BindDependency[usecase.IGetCampaignUnSignedCandidates, usecase.GetCampaignUnSignedCandidatesUseCase](container, nil)

	libConfig.BindDependency[usecase.ICampaignProgress, usecase.CampaignProgressUseCase](container, nil)

	libConfig.BindDependency[usecase.ICommitSpecificSigningInfo, usecase.CommitSpecificSigningInfoUseCase](container, nil)

	fmt.Println("Wiring dependencies successully.")
}

// func GetComponent[AbstractType](ctx iris.Context) {

// 	return ctx.Values().Get()
// }
