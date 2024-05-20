package main

import (
	"app/infrastructure/db"
	"app/internal/api"
	"app/internal/bootstrap"
	"app/internal/config"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofor-little/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	HOSTNAMES []string
)

func init() {

	cwd, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	env.Load(filepath.Join(cwd, ".env"))

	HOSTNAMES = bootstrap.RetrieveCORSHosts()

	db.GetDB()
}

func main() {

	var app *iris.Application = iris.New()

	app.Validator = validator.New()

	//app.UseGlobal(middleware.RedirectInternalError())
	app.UseRouter(
		cors.New(cors.Options{
			AllowedOrigins:   HOSTNAMES,
			AllowCredentials: true,
		}),
	)

	config.InitializeDatabase(app)
	config.RegisterServices(app)

	// registerServices(app)
	// registerDependencies(app)

	api.Init(app)
	app.Listen(
		env.Get("HTTP_PORT", ":80"),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithOptimizations,
	)
}

// func registerDependencies(app *iris.Application) {

// 	var container *hero.Container = app.ConfigureContainer().Container

// 	container.Register(loadDbClient())
// 	container.Register(db.GetDB())

// 	app.RegisterDependency(loadDbClient())
// 	// app.RegisterDependency(db.GetDB())
// }

// func registerServices(app *iris.Application) {

// 	//app.RegisterDependency(initAuthService())

// 	var container *hero.Container = app.ConfigureContainer().Container

// 	// container.Register(func(auth *authService.AuthenticateService) authService.IAuthService {

// 	// 	return auth
// 	// })

// 	auth := new(authService.AuthenticateService)
// 	Dep := container.Register(auth)
// 	Dep.DestType = libCommon.Wrap[authService.IAuthService]()

// }

// // func initAuthService() *authService.AuthenticateService {

// // 	s := new(authService.AuthenticateService)
// // 	//s.SetDB(db.GetDB())

// // 	return s
// // }

// func loadDbClient() *mongo.Client {

// 	client, err := db.GetClient()

// 	if err != nil {

// 		panic(err)
// 	}

// 	return client
// }
