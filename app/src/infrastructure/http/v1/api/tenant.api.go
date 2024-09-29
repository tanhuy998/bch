package api

import (
	"app/internal/controller"
	"app/internal/middleware"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func initTenantApi(app router.Party) {

	router := app.Party("tenants")

	container := router.ConfigureContainer().Container

	wrapper := mvc.New(router)

	wrapper.Router.Use(
		middleware.SecretAuth,
	)

	wrapper.Handle(
		new(controller.TenantController),
		// applyRoutes(func(activator *mvc.ControllerActivator) {

		// 	// activator.Handle(
		// 	// 	"POST", "/agent", "CreateTenantAgent",
		// 	// 	middleware.BindPresenters[requestPresenter.CreateTenantAgentRequest, responsePresenter.CreateTenantAgentResponse](container),
		// 	// )

		// 	activator.Handle(
		// 		"POST", "/", "CreateTenant",
		// 		middleware.BindPresenters[requestPresenter.CreateTenantRequest, responsePresenter.CreateTenantResponse](container),
		// 	)
		// }),
	)
}
