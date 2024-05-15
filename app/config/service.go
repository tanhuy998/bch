package config

import (
	"app/app/db"
	libCommon "app/app/lib/common"
	"app/app/repository"
	authService "app/app/service/auth"
	"fmt"
	"reflect"

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
	BindDependency[repository.ICampaignRepository](
		container, new(repository.CampaignRepository).Init(db),
	)
	BindDependency[repository.ICandidateRepository](
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
	BindDependency[authService.IAuthService, authService.AuthenticateService](container, nil)
	fmt.Println("Auth service initialized.")
}

func BindDependency[AbstractType any, ConcreteType any](container *hero.Container, concreteVal *ConcreteType) {

	abstractType := libCommon.Wrap[AbstractType]()

	if abstractType.Kind() != reflect.Interface {

		panic(
			fmt.Sprintf(
				"Could not use %s as abstract type which is not an interface",
				abstractType.String(),
			),
		)
	}

	if concreteVal == nil {

		concreteVal = new(ConcreteType)
	}

	//if !libCommon.GetOriginalTypeOf(concreteVal).Implements(abstractType) {
	if _, ok := any(concreteVal).(AbstractType); !ok {
		//if reflect.TypeOf(concreteVal).Implements(libCommon.InterfaceOf((*AbstractType)(nil))) {

		panic(
			fmt.Sprintf(
				"Could not bind concrete type %s as interface %s",
				reflect.TypeOf(concreteVal).String(),
				abstractType.String(),
			),
		)
	}

	dep := container.Register(concreteVal)
	dep.DestType = abstractType
}

// func GetComponent[AbstractType](ctx iris.Context) {

// 	return ctx.Values().Get()
// }
