package api

import (
	"app/app/controller"
	"app/app/middleware"
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCampaignGroupApi(app *iris.Application) {

	router := app.Party("/campaign")

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	})

	wrapper := mvc.New(router)
	wrapper.Handle(
		new(controller.CampaignController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			authField := authService.AuthorizationField("campaign")

			activator.Handle(
				"GET", "/{uuid:string}", "GetCampaign",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			activator.Handle(
				"POST", "/", "NewCampaign",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)

			activator.Handle(
				"PUT", "/{uuid:string}", "UpdateCampaign",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)

			activator.Handle(
				"DELETE", "/{uuid:string}", "DeleteCampaign",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)
		}),
	)
}
