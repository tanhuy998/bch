package config

import (
	libService "app/app/config/lib"
	"app/app/db"
	"app/app/repository"
	adminService "app/app/service/admin"
	authService "app/app/service/auth"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	AUTH         = "auth_service"
	DBMS_CLIENT  = "dbms_client"
	DB           = "db_instancce"
	REQUEST_BODY = "request_body"
)

func InitializeDatabase(app *iris.Application) {

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
	libService.BindDependency[repository.ICampaignRepository](
		container, new(repository.CampaignRepository).Init(db),
	)
	libService.BindDependency[repository.ICandidateRepository](
		container, new(repository.CandidateRepository).Init(db),
	)
	fmt.Println("Repositories Initialized.")
}

func RegisterServices(app *iris.Application) {

	var container *hero.Container = app.ConfigureContainer().Container

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
	libService.BindAndMapDependencyToContext[authService.IAuthService, authService.AuthenticateService](container, nil, AUTH)
	fmt.Println("Auth service initialized.")

	/*
		Bind Admin Campaign controller dependent services
	*/
	libService.BindDependency[adminService.IDeleteCampaign, adminService.AdminDeleteCampaignService](container, nil)
	libService.BindDependency[adminService.ILaunchNewCampaign, adminService.AdminLaunchNewCampaignService](container, nil)
	libService.BindDependency[adminService.IModifyExistingCampaign, adminService.AdminModifyExistingCampaign](container, nil)

	/*
		Bind Admin Candidate controller dependent services
	*/
	libService.BindDependency[adminService.IDeleteCandidate, adminService.AdminDeleteCandidateService](container, nil)
	libService.BindDependency[adminService.IAddNewCandidate, adminService.AdminAddNewCandidateToCampaign](container, nil)
	libService.BindDependency[adminService.IModifyCandidate, adminService.AdminModifyCandidate](container, nil)
}

// func GetComponent[AbstractType](ctx iris.Context) {

// 	return ctx.Values().Get()
// }
