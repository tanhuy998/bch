package restfull

import (
	"app/infrastructure/restfull/common/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type (
	RestFullAPIInitializationOption func(api iris.Party)
)

func configureLocalDependencies(app iris.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			api.EnableStructDependents()

			api.EnableStructDependents().RegisterDependency(
				new(middleware.AuthenticateMiddlewareHandler),
			)

			api.EnableStructDependents().RegisterDependency(
				new(middleware.InternalAccessLogHandler),
			).Explicitly()
		},
	)
}

func applyOptions(app *iris.Application, options ...RestFullAPIInitializationOption) {

	for _, optionFn := range options {

		if optionFn == nil {
			continue
		}

		optionFn(app)
	}
}
