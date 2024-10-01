package api

import (
	"app/src/infrastructure/http/api/v1/controller"
	"app/src/infrastructure/http/middleware"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

// func initCampaignGroup(app *flamego.Flame) {

// 	app.Group("/campaign", func() {

// 	})
// }

func initCampaignGroupApi(app router.Party) *mvc.Application {

	router := app.Party("/campaigns")

	container := router.ConfigureContainer().Container

	router.Use(
		middleware.Auth(
			container,
		),
	)

	controller := new(controller.CampaignController)

	// mvc.New(router).Handle(
	// 	controller,
	// 	applyRoutes(func(activator *mvc.ControllerActivator) {

	// 		activator.Handle(
	// 			"GET", "/", "GetCampaignListOnPage",
	// 			middleware.Authorize(authService.AuthorizationLicense{
	// 				Fields: campaignField,
	// 				Groups: []authService.AuthorizationGroup{auth_commander_group},
	// 			}),
	// 			middleware.BindRequest[requestPresenter.GetPendingCampaignRequest](),
	// 		).SetName("GET_LIST_CAMPAIGNS")
	// 	}),
	// ).EnableStructDependents()

	// router = router.Party("/")

	wrapper := mvc.New(router)
	return wrapper.Handle(
		controller,
		// applyRoutes(func(activator *mvc.ControllerActivator) {

		// 	//activator.Handle("GET", "/", "HealthCheck")

		// }),
	)
}
