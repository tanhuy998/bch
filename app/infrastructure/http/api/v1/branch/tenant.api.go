package api

import (
	"app/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func initTenantApi(app router.Party) *mvc.Application {

	router := app.Party("tenants")

	container := router.ConfigureContainer().Container

	wrapper := mvc.New(router)

	// wrapper.Router.Use(
	// 	middleware.SecretAuth,
	// )

	wrapper.Handle(
		new(controller.TenantController).BindDependencies(container),
	)

	return wrapper
}
