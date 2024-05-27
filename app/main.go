package main

import (
	"app/infrastructure/db"
	"app/internal/api"
	"app/internal/bootstrap"
	"app/internal/config"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	host_names           []string
	cors_allowed_methods []string = []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH"}
	server_ssl_cert string
	server_ssl_key  string
)

func init() {

	readSSlCert()

	godotenv.Load()

	host_names = bootstrap.RetrieveCORSHosts()

	err := db.CheckDBConnection()

	if err != nil {

		panic(err)
	}

	db.GetDB()
}

func main() {

	var app *iris.Application = iris.New()

	app.Validator = validator.New()

	// app.UseGlobal(func(ctx iris.Context) {

	// 	fmt.Println(ctx.Method(), ctx.RequestPath(false))

	// 	ctx.Next()
	// })

	// app.WrapRouter(methodoverride.New(
	// 	methodoverride.SaveOriginalMethod("_originalMethod"),
	// ))

	//app.UseGlobal(middleware.RedirectInternalError())
	app.UseRouter(
		cors.New(cors.Options{
			AllowedOrigins:   host_names,
			AllowedMethods:   cors_allowed_methods,
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
	// 				return next(ctx, val)"_originalMethod"
	// 			default:
	// 				fmt.Println(reflect.TypeOf(val).String())
	// 				return next(ctx, val)
	// 			}
	// 		}
	// 	})

	// app.Use(versioning.Aliases(versioning.AliasMap{
	// 	versioning.Empty: "1.0.0",
	// }))

	// v1 := versioning.NewGroup(app, ">=1.0.1 <2.0.0")

	config.InitializeDatabase(app)
	config.RegisterServices(app)

	// registerServices(app)
	// registerDependencies(app)

	api.Init(app)

	app.Run(
		iris.TLS(
			os.Getenv("HTTP_PORT"),
			server_ssl_cert,
			server_ssl_key,
		),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithOptimizations,
	)

	// app.Listen(
	// 	env.Get("HTTP_PORT", ":80"),
	// 	iris.WithoutBodyConsumptionOnUnmarshal,
	// 	iris.WithOptimizations,
	// )
}

func readSSlCert() {
	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	d, err := os.ReadFile(path.Join(__dir, "cert.pem"))

	if err != nil {

		panic(err)
	}

	server_ssl_cert = string(d)

	d, err = os.ReadFile(path.Join(__dir, "key.pem"))

	if err != nil {

		panic(err)
	}

	server_ssl_key = string(d)
}
