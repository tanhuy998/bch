package controller

import (
	"app/domain/presenter"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/refactor/infrastructure/http/middleware"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type SignController struct {
	CommitCandidateSigningInfoUseCase    usecase.ICommitCandidateSigningInfo
	GetSingleCandidateSigningInfoUseCase usecase.IGetSingleCandidateSigningInfo
	CheckSigningExistenceUseCase         usecase.ICheckSigningExistence
	CommitSpecificSigningInfoUseCase     usecase.ICommitSpecificSigningInfo
}

func (this *SignController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	/*
		Get Signing info of a candidate
	*/
	activator.Handle(
		"GET", "/candidate/{candidateUUID}", "GetSingleCandidateSigningInfo",
		middleware.BindPresenters[requestPresenter.GetSingleCandidateSigningInfoRequest, responsePresenter.GetSingleCandidateSigningInfoResponse](container),
	)

	/*
		Post signing info of a candidate
	*/
	activator.Handle(
		"PATCH", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "CommitCandidateSigningInfo",
		middleware.BindPresenters[requestPresenter.CommitCandidateSigningInfoRequest, responsePresenter.CommitCandidateSigningInfoResponse](container),
	)

	activator.Handle(
		"HEAD", "/pending/campaign/{campaignUUID}/candidate/{candidateUUID}", "CheckSigningExistence",
		middleware.BindPresenters[requestPresenter.CheckSigningExistenceRequest, presenter.IEmptyPresenter](container),
	)

	activator.Handle(
		"PUT", "/campaign/{campaignUUID}/candidate/{candidateUUID}", "CommitSpecificSigningInfo",
		middleware.BindPresenters[requestPresenter.CommitSpecificSigningInfo, responsePresenter.CommitSpecificSigningInfoResponse](container),
	)
}

func (this *SignController) GetSingleCandidateSigningInfo(
	input *requestPresenter.GetSingleCandidateSigningInfoRequest,
	output *responsePresenter.GetSingleCandidateSigningInfoResponse,
) (mvc.Result, error) {

	return this.GetSingleCandidateSigningInfoUseCase.Execute(input, output)
}

func (this *SignController) CommitCandidateSigningInfo(
	input *requestPresenter.CommitCandidateSigningInfoRequest,
	output *responsePresenter.CommitCandidateSigningInfoResponse,
) (mvc.Result, error) {

	return this.CommitCandidateSigningInfoUseCase.Execute(input, output)
}

func (this *SignController) CheckSigningExistence(
	input *requestPresenter.CheckSigningExistenceRequest,
) (mvc.Result, error) {

	return this.CheckSigningExistenceUseCase.Execute(input)
}

func (this *SignController) CommitSpecificSigningInfo(
	input *requestPresenter.CommitSpecificSigningInfo,
	output *responsePresenter.CommitSpecificSigningInfoResponse,
) (mvc.Result, error) {

	return this.CommitSpecificSigningInfoUseCase.Execute(input, output)
}
