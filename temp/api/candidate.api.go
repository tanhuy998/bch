package api

import (
	"app/src/infrastructure/http/api/v1/controller"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func initCandidateGroupApi(app router.Party) *mvc.Application {

	router := app.Party("/candidates")

	// container := router.ConfigureContainer(func(api *iris.APIContainer) {
	// 	/*
	// 		for dependencies injection
	// 	*/
	// 	api.Use(middleware.Authentication())
	// }).Container

	wrapper := mvc.New(router)
	return wrapper.Handle(
		new(controller.CandidateController),
		// applyRoutes(func(activator *mvc.ControllerActivator) {

		// 	// 			/*
		// 	// 				Get Candidate list in pagination form of a specific campaign
		// 	// 			*/
		// 	// 			activator.Handle(
		// 	// 				"GET", "/campaign/{campaignUUID}", "GetCampaignCandidateList",

		// 	// 				middleware.BindPresenters[requestPresenter.GetCampaignCandidateListRequest, responsePresenter.GetCampaignCandidateListResponse](container),
		// 	// 			).SetName("GET_CANDIDATE_LIST_OF_EXISTING_CAMPAIGN")

		// 	// 			/*
		// 	// 				Seach Candidates by given infomations
		// 	// 			*/
		// 	// 			activator.Handle(
		// 	// 				"GET", "/search", "SearchByInformation",
		// 	// \
		// 	// 			)

		// 	// 			/*
		// 	// 				Get a specific candidate by uuid
		// 	// 			*/
		// 	// 			activator.Handle(
		// 	// 				"GET", "/{uuid}", "GetSingleCandidateByUUID",

		// 	// 				middleware.BindPresenters[requestPresenter.GetSingleCandidateRequest, responsePresenter.GetSingleCandidateResponse](container),
		// 	// 			).SetName("GET_SINGLE_CANDIDATE_BY_UUID")

		// 	// 			activator.Handle(
		// 	// 				"GET", "/signed/campaign/{campaignUUID}", "GetSignedCandidates",
		// 	// 				middleware.BindPresenters[requestPresenter.GetCampaignSignedCandidatesRequest, responsePresenter.GetCampaignSignedCandidatesResponse](container),
		// 	// 			).SetName("GET_CAMPAIGN_SIGNED_CANDIDATES")

		// 	// 			activator.Handle(
		// 	// 				"GET", "/unsigned/campaign/{campaignUUID}", "GetUnSignedCandidates",
		// 	// 				middleware.BindPresenters[requestPresenter.GetCampaignUnSignedCandidates, responsePresenter.GetCampaignUnSignedCandidates](container),
		// 	// 			).SetName("GET_CAMPAIGN_UNSIGNED_CANDIDATES")

		// 	// 			/*
		// 	// 				Post a candidate to a specific campaign
		// 	// 			*/
		// 	// 			activator.Handle(
		// 	// 				"POST", "/campaign/{campaignUUID}", "PostCandidate",

		// 	// 				middleware.BindPresenters[requestPresenter.AddCandidateRequest, responsePresenter.AddNewCandidateResponse](container),
		// 	// 			).SetName("ADD_NEW_CANDIDATE_TO_EXISTING_CAMPAIGN")

		// 	// 			/*
		// 	// 				Update information of a candidate
		// 	// 			*/
		// 	// 			activator.Handle(
		// 	// 				"PATCH", "/{uuid}", "UpdateCandidate",

		// 	// 				middleware.BindPresenters[requestPresenter.ModifyExistingCandidateRequest, responsePresenter.ModifyExistingCandidateResponse](container),
		// 	// 			).SetName("MODIFY_EXISTING_CANDIDATE")

		// 	// activator.Handle(
		// 	// 	"PATCH", "/detail/{uuid}", "UpdateCandidateDetailInfo",
		// 	// 	middleware.Authorize(
		// 	// 		authService.AuthorizationLicense{
		// 	// 			Fields: candidateField,
		// 	// 			Groups: []authService.AuthorizationGroup{auth_commander_group, auth_member_group},
		// 	// 		},
		// 	// 	),
		// 	// )

		// 	/*
		// 		Delete a
		// 	*/
		// 	// activator.Handle(
		// 	// 	"DELETE", "/{uuid}", "DeleteCandidate",
		// 	// 	middleware.Authorize(authService.AuthorizationLicense{
		// 	// 		Fields: candidateField,
		// 	// 		//Groups: []authService.AuthorizationGroup{auth_commander_group},
		// 	// 		Claims: []authService.AuthorizationClaim{auth_delete_claim},
		// 	// 	}),
		// 	// )
		// }),
	)
}
