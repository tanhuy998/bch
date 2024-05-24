package api

import (
	requestPresenter "app/domain/presenter/request"
	"app/internal/controller"
	"app/internal/middleware"
	authService "app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// func initCampaignGroup(app *flamego.Flame) {

// 	app.Group("/campaign", func() {

// 	})
// }

func initCampaignGroupApi(app *iris.Application) *mvc.Application {

	router := app.Party("/campaigns")

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	})

	controller := new(controller.CampaignController)
	campaignField := authService.AuthorizationField("campaigns")

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
		applyRoutes(func(activator *mvc.ControllerActivator) {

			//activator.Handle("GET", "/", "HealthCheck")

			activator.Handle(
				"GET", "/", "GetCampaignListOnPage",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequest[requestPresenter.GetCampaignListRequest](),
			).SetName("GET_LIST_CAMPAIGNS")

			activator.Handle(
				"GET", "/{uuid:uuid}", "GetCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
				middleware.BindRequest[requestPresenter.GetSingleCampaignRequest](),
			).SetName("GET_SINGLE_CAMPAIGN")

			activator.Handle(
				"GET", "/pending", "GetPendingCampaigns",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
				}),
				middleware.BindRequest[requestPresenter.GetPendingCampaignRequest](),
			).SetName("GET_PENDING_CAMPAIGNS")

			activator.Handle(
				"POST", "/", "NewCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Claims: []authService.AuthorizationClaim{auth_post_claim},
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequest[requestPresenter.LaunchNewCampaignRequest](),
			).SetName("LAUNCH_NEW_CAMPAIGN")

			activator.Handle(
				"PATCH", "/{uuid:string}", "UpdateCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequest[requestPresenter.UpdateCampaignRequest](),
			).SetName("UPDATE_CAMPAIGN")

			activator.Handle(
				"DELETE", "/{uuid:string}", "DeleteCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Claims: []authService.AuthorizationClaim{auth_post_claim},
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequest[requestPresenter.DeleteCampaignRequest](),
			).SetName("DELETE_CAMPAIGN")
		}),
	)
}
