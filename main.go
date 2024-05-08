package main

import (
	"app/app/api"
	"app/app/config"
	"app/app/db"
	"app/app/middleware"
	authService "app/app/service/auth"

	"github.com/go-playground/validator/v10"
	"github.com/gofor-little/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	loadConfig()

	var app *iris.Application = iris.New()

	app.Validator = validator.New()

	app.UseGlobal(middleware.RedirectInternalError())
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

func loadConfig() {

	config.InitEnv()
}

func registerDependencies(app *iris.Application) {

	app.RegisterDependency(loadDbClient())
	app.RegisterDependency(db.GetDB())
}

func registerServices(app *iris.Application) {

	app.RegisterDependency(initAuthService())
}

func initAuthService() *authService.AuthenticateService {

	s := new(authService.AuthenticateService)
	//s.SetDB(db.GetDB())

	return s
}

func loadDbClient() *mongo.Client {

	client, err := db.GetClient()

	if err != nil {

		panic(err)
	}

	return client
}
