package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/src/infrastructure/http/middleware"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type CandidateController struct {
	AddNewCandidateUseCase               usecase.IAddNewCandidate
	ModifyCandidateUseCase               usecase.IModifyExistingCandidate
	DeleteCandidateUseCase               usecase.IDeleteCandidate
	GetCampaignCandidateListUseCase      usecase.IGetCampaignCandidateList
	GetSingleCandidateByUUIDUseCase      usecase.IGetSingleCandidateByUUID
	GetCampaignSignedCandidatesUseCase   usecase.IGetCampaignSignedCandidates
	GetCampaignUnSignedCandidatesUseCase usecase.IGetCampaignUnSignedCandidates
}

func (this *CandidateController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	/*
				Get Candidate list in pagination form of a specific campaign
			*/
			activator.Handle(
				"GET", "/campaign/{campaignUUID}", "GetCampaignCandidateList",
				
				middleware.BindPresenters[requestPresenter.GetCampaignCandidateListRequest, responsePresenter.GetCampaignCandidateListResponse](container),
			).SetName("GET_CANDIDATE_LIST_OF_EXISTING_CAMPAIGN")

			/*
				Seach Candidates by given infomations
			*/
			activator.Handle(
				"GET", "/search", "SearchByInformation",
\
			)

			/*
				Get a specific candidate by uuid
			*/
			activator.Handle(
				"GET", "/{uuid}", "GetSingleCandidateByUUID",

				middleware.BindPresenters[requestPresenter.GetSingleCandidateRequest, responsePresenter.GetSingleCandidateResponse](container),
			).SetName("GET_SINGLE_CANDIDATE_BY_UUID")

			activator.Handle(
				"GET", "/signed/campaign/{campaignUUID}", "GetSignedCandidates",
				middleware.BindPresenters[requestPresenter.GetCampaignSignedCandidatesRequest, responsePresenter.GetCampaignSignedCandidatesResponse](container),
			).SetName("GET_CAMPAIGN_SIGNED_CANDIDATES")

			activator.Handle(
				"GET", "/unsigned/campaign/{campaignUUID}", "GetUnSignedCandidates",
				middleware.BindPresenters[requestPresenter.GetCampaignUnSignedCandidates, responsePresenter.GetCampaignUnSignedCandidates](container),
			).SetName("GET_CAMPAIGN_UNSIGNED_CANDIDATES")

			/*
				Post a candidate to a specific campaign
			*/
			activator.Handle(
				"POST", "/campaign/{campaignUUID}", "PostCandidate",

				middleware.BindPresenters[requestPresenter.AddCandidateRequest, responsePresenter.AddNewCandidateResponse](container),
			).SetName("ADD_NEW_CANDIDATE_TO_EXISTING_CAMPAIGN")

			/*
				Update information of a candidate
			*/
			activator.Handle(
				"PATCH", "/{uuid}", "UpdateCandidate",

				middleware.BindPresenters[requestPresenter.ModifyExistingCandidateRequest, responsePresenter.ModifyExistingCandidateResponse](container),
			).SetName("MODIFY_EXISTING_CANDIDATE")	
}

func (this *CandidateController) GetSingleCandidateByUUID(
	input *requestPresenter.GetSingleCandidateRequest,
	output *responsePresenter.GetSingleCandidateResponse,
) (mvc.Result, error) {

	return this.GetSingleCandidateByUUIDUseCase.Execute(input, output)
}

func (this *CandidateController) UpdateCandidateDetailInfo() {

}

func (this *CandidateController) PostCandidate(
	input *requestPresenter.AddCandidateRequest,
	output *responsePresenter.AddNewCandidateResponse,
) (mvc.Result, error) {

	return this.AddNewCandidateUseCase.Execute(input, output)
}

func (this *CandidateController) UpdateCandidate(
	input *requestPresenter.ModifyExistingCandidateRequest,
	output *responsePresenter.ModifyExistingCandidateResponse,
) (mvc.Result, error) {

	return this.ModifyCandidateUseCase.Execute(input, output)
}

func (this *CandidateController) DeleteCandidate(
	input *requestPresenter.DeleteCandidateRequest,
	output *responsePresenter.DeleteCandidateResponse,
) (mvc.Result, error) {

	return this.DeleteCandidateUseCase.Execute(input, output)
}

func (this *CandidateController) GetCampaignCandidateList(
	input *requestPresenter.GetCampaignCandidateListRequest,
	output *responsePresenter.GetCampaignCandidateListResponse,
) (mvc.Result, error) {

	return this.GetCampaignCandidateListUseCase.Execute(input, output)
}

func (this *CandidateController) SearchByInformation() {

}

func (this *CandidateController) GetSignedCandidates(
	input *requestPresenter.GetCampaignSignedCandidatesRequest,
	output *responsePresenter.GetCampaignSignedCandidatesResponse,
) (mvc.Result, error) {

	return this.GetCampaignSignedCandidatesUseCase.Execute(input, output)
}

func (this *CandidateController) GetUnSignedCandidates(
	input *requestPresenter.GetCampaignUnSignedCandidates,
	output *responsePresenter.GetCampaignUnSignedCandidates,
) (mvc.Result, error) {

	return this.GetCampaignUnSignedCandidatesUseCase.Execute(input, output)
}
