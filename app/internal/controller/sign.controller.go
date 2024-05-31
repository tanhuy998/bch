package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	usecase "app/useCase"

	"github.com/kataras/iris/v12/mvc"
)

type SignController struct {
	CommitCandidateSigningInfoUseCase usecase.ICommitCandidateSigningInfo
}

func (this *SignController) GetSigningInfo() {

}

func (this *SignController) CommitCandidateSigningInfo(
	input *requestPresenter.CommitCandidateSigningInfoRequest,
	output *responsePresenter.CommitCandidateSigningInfoResponse,
) (mvc.Result, error) {

	return this.CommitCandidateSigningInfoUseCase.Execute(input, output)
}
