package config

import (
	"app/infrastructure/db"
	libConfig "app/lib/config"
	"app/repository"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	authService "app/service/auth"
	candidateService "app/service/candidate"
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
	libConfig.BindDependency[repository.ICampaignRepository](
		container, new(repository.CampaignRepository).Init(db),
	)
	libConfig.BindDependency[repository.ICandidateRepository](
		container, new(repository.CandidateRepository).Init(db),
	)
	fmt.Println("Repositories Initialized.")
}

func RegisterServices(app router.Party) {

	var container *hero.Container = app.ConfigureContainer().EnableStructDependents().Container

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

	libConfig.BindAndMapDependencyToContext[authService.IAuthService, authService.AuthenticateService](container, nil, AUTH)

	fmt.Println("Auth service initialized.")

	//libConfig.BindDependency[port.IActionResult, usecase.ActionResultUseCase](container, nil)

	/*
		init app validator
	*/
	libConfig.BindDependency[context.Validator, validator.Validate](container, validator.New())
	libConfig.BindDependency[actionResultService.IActionResult, actionResultService.ResponseResultService](container, nil)
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

	libConfig.BindDependency[candidateService.ICommitCandidateSigningInfo, candidateService.CommitCandidateSigningInfoService](container, nil)
	libConfig.BindDependency[candidateService.IGetSingleCandidateSigningInfo, candidateService.GetSingleCandidateSigningInfoService](container, nil)

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
}

// func GetComponent[AbstractType](ctx iris.Context) {

// 	return ctx.Values().Get()
// }
