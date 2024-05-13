package main

import (
	"app/app/api"
	"app/app/config"
	"app/app/db"
	libCommon "app/app/lib/common"
	authService "app/app/service/auth"

	"github.com/go-playground/validator/v10"
	"github.com/gofor-little/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	loadConfig()

	var app *iris.Application = iris.New()

	app.Validator = validator.New()

	//app.UseGlobal(middleware.RedirectInternalError())
	app.UseRouter(
		cors.New(cors.Options{
			AllowedOrigins:   config.RetrieveCORSHosts(),
			AllowCredentials: true,
		}),
	)

	registerServices(app)
	registerDependencies(app)

	api.Init(app)
	app.Listen(env.Get("HTTP_PORT", ":80"))
}

// func initIrisApp() *iris.Application {

// }

func loadConfig() {

	config.InitEnv()
}

func registerDependencies(app *iris.Application) {

	var container *hero.Container = app.ConfigureContainer().Container

	container.Register(loadDbClient())
	container.Register(db.GetDB())

	app.RegisterDependency(loadDbClient())
	// app.RegisterDependency(db.GetDB())
}

func registerServices(app *iris.Application) {

	//app.RegisterDependency(initAuthService())

	var container *hero.Container = app.ConfigureContainer().Container

	// container.Register(func(auth *authService.AuthenticateService) authService.IAuthService {

	// 	return auth
	// })

	auth := new(authService.AuthenticateService)
	Dep := container.Register(auth)
	Dep.DestType = libCommon.InterfaceOf((*authService.IAuthService)(nil))
}

// func initAuthService() *authService.AuthenticateService {

// 	s := new(authService.AuthenticateService)
// 	//s.SetDB(db.GetDB())

// 	return s
// }

func loadDbClient() *mongo.Client {

	client, err := db.GetClient()

	if err != nil {

		panic(err)
	}

	return client
}
