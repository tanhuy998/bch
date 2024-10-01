package usecase

import (
	actionResultService "app/service/actionResult"
	candidateService "app/service/candidate"
	requestPresenter "app/src/presenter/request"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICheckSigningExistence interface {
		Execute(input *requestPresenter.CheckSigningExistenceRequest) (mvc.Result, error)
	}

	CheckSigningExistenceUseCase struct {
		CheckSigningExitenceService candidateService.ICheckSigningExistence
		ActionResultService         actionResultService.IActionResult
	}
)

func (this *CheckSigningExistenceUseCase) Execute(
	input *requestPresenter.CheckSigningExistenceRequest,
) (mvc.Result, error) {

	exist, err := this.CheckSigningExitenceService.Serve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return nil, err
	}

	var statusCode int

	if exist {

		statusCode = 204
	} else {

		statusCode = 404
	}

	return this.ActionResultService.Prepare().SetCode(statusCode), nil
}
