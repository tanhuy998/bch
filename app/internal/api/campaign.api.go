package api

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
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

	container := router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	}).Container

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
				middleware.BindPresenters[requestPresenter.GetCampaignListRequest, responsePresenter.GetCampaignListResponse](container),
			).SetName("GET_LIST_CAMPAIGNS")

			activator.Handle(
				"GET", "/{uuid:uuid}", "GetCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
				middleware.BindPresenters[requestPresenter.GetSingleCampaignRequest, responsePresenter.GetSingleCampaignResponse](container),
			).SetName("GET_SINGLE_CAMPAIGN")

			activator.Handle(
				"GET", "/{uuid:uuid}/progress", "GetCampaignProgress",
				middleware.BindPresenters[requestPresenter.CampaignProgressRequestPresenter, responsePresenter.CampaignProgressResponsePresenter](container),
			)

			activator.Handle(
				"GET", "/pending", "GetPendingCampaigns",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
				}),
				middleware.BindPresenters[requestPresenter.GetPendingCampaignRequest, responsePresenter.GetPendingCampaingsResponse](container),
			).SetName("GET_PENDING_CAMPAIGNS")

			activator.Handle(
				"POST", "/", "NewCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Claims: []authService.AuthorizationClaim{auth_post_claim},
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindPresenters[requestPresenter.LaunchNewCampaignRequest, responsePresenter.LaunchNewCampaignResponse](container),
			).SetName("LAUNCH_NEW_CAMPAIGN")

			activator.Handle(
				"PATCH", "/{uuid:uuid}", "UpdateCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindPresenters[requestPresenter.UpdateCampaignRequest, responsePresenter.UpdateCampaignResponse](container),
			).SetName("UPDATE_CAMPAIGN")

			// activator.Handle(
			// 	"DELETE", "/{uuid:uuid}", "DeleteCampaign",
			// 	middleware.Authorize(authService.AuthorizationLicense{
			// 		Fields: campaignField,
			// 		Claims: []authService.AuthorizationClaim{auth_post_claim},
			// 		//Groups: []authService.AuthorizationGroup{auth_commander_group},
			// 	}),
			// 	middleware.BindPresenters[requestPresenter.DeleteCampaignRequest, responsePresenter.DeleteCampaignResponse](container),
			// ).SetName("DELETE_CAMPAIGN")

			activator.Handle(
				"PATCH", "/test/{uuid:uuid}", "TestPatch",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindPresenters[requestPresenter.UpdateCampaignRequest, responsePresenter.UpdateCampaignResponse](container),
			)
		}),
	)
}
