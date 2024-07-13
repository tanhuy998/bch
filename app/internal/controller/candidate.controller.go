package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type CandidateController struct {
	AddNewCandidateUseCase             usecase.IAddNewCandidate
	ModifyCandidateUseCase             usecase.IModifyExistingCandidate
	DeleteCandidateUseCase             usecase.IDeleteCandidate
	GetCampaignCandidateListUseCase    usecase.IGetCampaignCandidateList
	GetSingleCandidateByUUIDUseCase    usecase.IGetSingleCandidateByUUID
	GetCampaignSignedCandidatesUseCase usecase.IGetCampaignSignedCandidates
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
