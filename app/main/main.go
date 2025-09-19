package main

import (
	irisConfig "app/infrastructure/http/api/v1/config"
	"app/infrastructure/restfull"
	"app/internal/bootstrap"
	"app/internal/rpc"
	"app/internal/watcher"

	"app/main/internal/ioc"
	"app/main/internal/manifest"
	"app/main/internal/tls"

	"app/main/internal/dependencies/log"

	"github.com/gofor-little/env"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

func init() {

	bootstrap.Boot()
}

func main() {

	var restfullAPI *iris.Application = restfull.NewAPI(
		func(api iris.Party) {

			ioc.RegisterServices(api)
		},
	)


	restfullAPI.UseGlobal(
		manifest.JustedSources,
	)

	// for _, v := range restfullAPI.GetRoutes() {

	// 	fmt.Println(v)
	// }

	// restfullAPI.Configure(

	// 	iris.WithoutBodyConsumptionOnUnmarshal,
	// 	iris.WithOptimizations,
	// )
	// globalContainer := app.ConfigureContainer().EnableStructDependents().Container

	watcher.Watch(
		func() {

			defer irisConfig.ConfigureLogger(restfullAPI).Close()

			restfullAPI.Run(
				iris.TLS(
					env.Get("HTTP_PORT", ":443"),
					tls.GetSSLCert(),
					tls.GetSSLKey(),
				),
				iris.WithHostProxyHeader(
					"Host",
					"X-Real-IP",
					"X-Forwarded-For",
					"X-Forwarded-Proto",
				),
				iris.WithoutBodyConsumptionOnUnmarshal,
				iris.WithOptimizations,
			)

			log.Main().Println("Http server closed.")
		},
		func() {

			rpc.Listen()
			log.Main().Println("Rpc server closed.")
		},
	)

	watcher.Wait()
}
