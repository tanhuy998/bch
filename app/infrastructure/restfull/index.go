package restfull

import (
	v1 "app/infrastructure/restfull/api/v1"
	"app/infrastructure/restfull/common/log"
	"app/infrastructure/restfull/common/middleware"
	libCommon "app/internal/lib/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type (
	RestFullAPIInitializationOption func(api iris.Party)
)

func NewAPI(options ...RestFullAPIInitializationOption) *iris.Application {

	defer libCommon.LMessureTime("Restfull API initialization time", log.Logger())()

	app := iris.New()

	applyOptions(app, options...)

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			api.EnableStructDependents().RegisterDependency(
				new(middleware.AuthenticateMiddlewareHandler),
			)
		},
	)

	app.UseRouter(
		middleware.InternalAccessLog(
			app.ConfigureContainer().Container,
		),
	)

	v1.Initialize(app)

	return app
}

func applyOptions(app *iris.Application, options ...RestFullAPIInitializationOption) {

	for _, optionFn := range options {

		optionFn(app)
	}
}
