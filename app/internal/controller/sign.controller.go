package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type SignController struct {
	CommitCandidateSigningInfoUseCase    usecase.ICommitCandidateSigningInfo
	GetSingleCandidateSigningInfoUseCase usecase.IGetSingleCandidateSigningInfo
	CheckSigningExistenceUseCase         usecase.ICheckSigningExistence
	CommitSpecificSigningInfoUseCase     usecase.ICommitSpecificSigningInfo
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
