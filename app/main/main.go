package main

import (
	irisConfig "app/infrastructure/http/api/v1/config"
	"app/infrastructure/restfull"
	v1 "app/infrastructure/restfull/api/v1"
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

	var restfullAPI *restfull.API = restfull.NewAPI(
		ioc.RegisterServices,
	).Build(v1.Initialize)

	restfullAPI.UseGlobal(
		manifest.JustedSources,
	)

	watcher.Watch(
		func() {

			defer irisConfig.ConfigureLogger(restfullAPI.Application).Close()

			err := restfullAPI.Run(
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

			log.Main().Println("Http server closed.", err)
		},
		func() {

			rpc.Listen()
			log.Main().Println("Rpc server closed.")
		},
	)

	watcher.Wait()
}
