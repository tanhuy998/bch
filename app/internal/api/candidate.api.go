package api

import (
	"app/domain/model"
	"app/internal/controller"
	"app/internal/middleware"
	authService "app/service/auth"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateGroupApi(app *iris.Application) *mvc.Application {

	router := app.Party("/candidates")

	router.ConfigureContainer(func(api *iris.APIContainer) {
		/*
			for dependencies injection
		*/
		api.Use(middleware.Authentication())
	})

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.CandidateController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			candidateField := authService.AuthorizationField("candidate")
			campaignField := authService.AuthorizationField("campaign")

			/*
				Get Candidate list in pagination form of a specific campaign
			*/
			activator.Handle(
				"GET", "/campaign/{uuid:string}", "GetCandidateByPage",
				middleware.Authorize(
					authService.AuthorizationLicense{
						Fields: candidateField,
						Claims: []authService.AuthorizationClaim{auth_get_claim},
					},
					authService.AuthorizationLicense{
						Fields: campaignField,
						Claims: []authService.AuthorizationClaim{auth_get_claim},
					},
				),
			)

			/*
				Seach Candidates by given infomations
			*/
			activator.Handle(
				"GET", "/search", "SearchByInformation",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: candidateField,
					Claims: []authService.AuthorizationClaim{authService.AuthorizationClaim("discover_claim")},
				}),
			)

			/*
				Get a specific candidate by uuid
			*/
			activator.Handle(
				"GET", "/{uuid:string}", "GetCandidate",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: candidateField,
					Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
			)

			/*
				Post a candidate to a specific campaign
			*/
			activator.Handle(
				"POST", "/campaign/{uuid:string}", "PostCandidate",
				middleware.Authorize(
					authService.AuthorizationLicense{
						Fields: candidateField,
						//Groups: []authService.AuthorizationGroup{auth_commander_group},
						Claims: []authService.AuthorizationClaim{auth_post_claim},
					},
					authService.AuthorizationLicense{
						Fields: campaignField,
						Claims: []authService.AuthorizationClaim{auth_post_claim},
					},
				),
				middleware.BindRequestBody[model.Candidate](),
			)

			/*
				Update information of a candidate
			*/
			activator.Handle(
				"PATCH", "/{uuid:string}", "UpdateCandidate",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: candidateField,
					Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
				}),
				middleware.BindRequestBody[model.Candidate](),
			)

			activator.Handle(
				"PATCH", "/detail/{uuid:string}", "UpdateCandidateDetailInfo",
				middleware.Authorize(
					authService.AuthorizationLicense{
						Fields: candidateField,
						Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
					},
				),
			)

			/*
				Delete a
			*/
			activator.Handle(
				"DELETE", "/{uuid:string}", "DeleteCandidate",
				middleware.Authorize(authService.AuthorizationLicense{
					Fields: candidateField,
					//Groups: []authService.AuthorizationGroup{auth_commander_group},
					Claims: []authService.AuthorizationClaim{auth_delete_claim},
				}),
			)
		}),
	)
}
