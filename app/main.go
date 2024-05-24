package main

import (
	"app/infrastructure/db"
	"app/internal/api"
	"app/internal/bootstrap"
	"app/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofor-little/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	HOSTNAMES []string
)

func init() {

	// cwd, err := os.Getwd()

	// if err != nil {

	// 	panic(err)
	// }

	// env.Load(filepath.Join(cwd, ".env"))

	godotenv.Load()

	HOSTNAMES = bootstrap.RetrieveCORSHosts()

	err := db.CheckDBConnection()

	if err != nil {

		panic(err)
	}

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

	defer config.ConfigureLogger(app).Close()

	// app.ConfigureContainer().
	// 	UseResultHandler(func(next iris.ResultHandler) iris.ResultHandler {
	// 		return func(ctx iris.Context, v interface{}) error {
	// 			fmt.Println("error catcher")
	// 			switch val := v.(type) {
	// 			case error:
	// 				fmt.Println("err")
	// 				return next(ctx, val)
	// 			case *mvc.Response:
	// 				fmt.Println("err res")
	// 				return next(ctx, val)
	// 			default:
	// 				fmt.Println(reflect.TypeOf(val).String())
	// 				return next(ctx, val)
	// 			}
	// 		}
	// 	})

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
