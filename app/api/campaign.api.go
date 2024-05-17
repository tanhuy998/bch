package api

import (
	"app/app/controller"
	"app/app/middleware"
	"app/app/model"
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// func initCampaignGroup(app *flamego.Flame) {

// 	app.Group("/campaign", func() {

// 	})
// }

func initCampaignGroupApi(app *iris.Application) {

	router := app.Party("/campaign")

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	})

	wrapper := mvc.New(router)
	wrapper.Handle(
		new(controller.CampaignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			campaignField := authService.AuthorizationField("campaign")

			//activator.Handle("GET", "/", "HealthCheck")

			activator.Handle(
				"GET", "/{uuid:string}", "GetCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			activator.Handle(
				"Get", "/page/{page:int}", "GetCampaignListOnPage",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			activator.Handle(
				"POST", "/", "NewCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Claims: []authService.AuthorizationClaim{auth_post_claim},
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequestBody[model.Campaign](),
			)

			activator.Handle(
				"PATCH", "/{uuid:string}", "UpdateCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
				middleware.BindRequestBody[model.Campaign](),
			)

			activator.Handle(
				"DELETE", "/{uuid:string}", "DeleteCampaign",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: campaignField,
					Claims: []authService.AuthorizationClaim{auth_post_claim},
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)
		}),
	)
}
