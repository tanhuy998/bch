package v1

import (
	api "app/infrastructure/http/api/v1/branch"
	"app/infrastructure/http/api/v1/config"

	"github.com/kataras/iris/v12"
)

func Initialize(app *iris.Application) {

	config.RegisterServices(app)

	api.Init(app)
}
