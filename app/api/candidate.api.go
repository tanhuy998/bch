package api

import (
	"app/app/controller"
	"app/app/middleware"
	authService "app/app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateGroupApi(app *iris.Application) {

	router := app.Party("/candidate")

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
	})

	wrapper := mvc.New(router)
	wrapper.Handle(
		new(controller.CandidateController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			authField := authService.AuthorizationField("candidate")

			activator.Handle(
				"GET", "/{uuid:string}", "GetCandidate",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			activator.Handle(
				"POST", "/", "PostCandidate",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)

			activator.Handle(
				"PUT", "/{uuid:string}", "UpdateCandidate",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			activator.Handle(
				"DELETE", "/{uuid:string}", "DeleteCandidate",
				middleware.Authorization(middleware.AuthorizationLicense{
					Fields: []authService.AuthorizationField{authField},
					Groups: []authService.AuthorizationGroup{auth_commander_group},
				}),
			)
		}),
	)
}
