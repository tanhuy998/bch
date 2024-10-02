package main

import (
	v1 "app/infrastructure/http"
	"app/infrastructure/http/api/v1/config"
	"app/internal/bootstrap"
	"app/internal/db"

	"os"
	"path"

	"github.com/gofor-little/env"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	//host_names           []string
	cors_allowed_methods []string = []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH",
	}
	cors_allowed_headers []string = []string{
		"Authorization",
	}
	server_ssl_cert string
	server_ssl_key  string
)

func init() {

	//godotenv.Load()

	db.Init()
	readSSlCert()

	//host_names = bootstrap.RetrieveCORSHosts()
	//config.InitializeAuthEncryptionData()

	err := db.CheckDBConnection()

	if err != nil {

		panic(err)
	}

	db := db.GetDB()

	config.InitDomainIndexes(db)
}

func main() {

	var app *iris.Application = iris.New()

	//app.Validator = validator.New()

	// app.WrapRouter(methodoverride.New(
	// 	methodoverride.SaveOriginalMethod("_originalMethod"),
	// ))

	//app.UseGlobal(middleware.RedirectInternalError())
	app.UseRouter(
		cors.New(cors.Options{
			AllowedOrigins:   bootstrap.GetHostNames(),
			AllowedMethods:   cors_allowed_methods,
			AllowedHeaders:   cors_allowed_headers,
			AllowCredentials: true,
		}),
	)

	defer config.ConfigureLogger(app).Close()

	// defer config.ConfigureLogger(app).Close()

	// // app.ConfigureContainer().
	// // 	UseResultHandler(func(next iris.ResultHandler) iris.ResultHandler {
	// // 		return func(ctx iris.Context, v interface{}) error {
	// // 			fmt.Println("error catcher")
	// // 			switch val := v.(type) {
	// // 			case error:
	// // 				fmt.Println("err")
	// // 				return next(ctx, val)
	// // 			case *mvc.Response:
	// // 				fmt.Println("err res")
	// // 				return next(ctx, val)"_originalMethod"
	// // 			default:
	// // 				fmt.Println(reflect.TypeOf(val).String())
	// // 				return next(ctx, val)
	// // 			}
	// // 		}
	// // 	})

	// // app.Use(versioning.Aliases(versioning.AliasMap{
	// // 	versioning.Empty: "1.0.0",
	// // }))

	// // v1 := versioning.NewGroup(app, ">=1.0.1 <2.0.0")

	// config.InitializeDatabase(app)
	// config.RegisterServices(app)

	// // registerServices(app)
	// // registerDependencies(app)

	// api.Init(app)

	// // app.Run(
	// // 	iris.TLS(
	// // 		os.Getenv("HTTP_PORT"),
	// // 		server_ssl_cert,
	// // 		server_ssl_key,
	// // 	),
	// // 	iris.WithoutBodyConsumptionOnUnmarshal,
	// // 	iris.WithOptimizations,
	// // )

	v1.Initialize(app)

	app.Listen(
		env.Get("HTTP_PORT", ":80"),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithOptimizations,
	)
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
