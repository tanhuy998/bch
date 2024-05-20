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

	router := app.Party("/campaign")

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	})

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.CampaignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			campaignField := authService.AuthorizationField("campaigns")

			//activator.Handle("GET", "/", "HealthCheck")

			activator.Handle(
				"GET", "/{uuid:string}", "GetCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
				//middleware.BindRequest[]()
			).SetName("GET_SINGLE_CAMPAIGN")

			activator.Handle(
				"Get", "/", "GetCampaignListOnPage",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequest[requestPresenter.GetPendingCampaignRequest](),
			).SetName("GET_LIST_CAMPAIGNS")

			activator.Handle(
				"GET", "/pending", "GetPendingCampaigns",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
				}),
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
