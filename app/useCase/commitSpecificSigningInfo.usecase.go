package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	"app/service/signingService"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICommitSpecificSigningInfo interface {
		Execute(
			input *requestPresenter.CommitSpecificSigningInfo,
			output *responsePresenter.CommitSpecificSigningInfoResponse,
		) (mvc.Result, error)
	}

	CommitSpecificSigningInfoUseCase struct {
		CommitSpecificSigningInfoService signingService.ICommitSpecificSigningInfo
		ActionResult                     actionResultService.IActionResult
	}
)

func (this *CommitSpecificSigningInfoUseCase) Execute(
	input *requestPresenter.CommitSpecificSigningInfo,
	output *responsePresenter.CommitSpecificSigningInfoResponse,
) (mvc.Result, error) {

	err := this.CommitSpecificSigningInfoService.Serve(input.CandidateUUID, input.Data)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResult.ServeResponse(output)
}
