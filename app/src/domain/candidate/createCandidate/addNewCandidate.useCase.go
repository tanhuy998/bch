package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	"time"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IAddNewCandidate interface {
		Execute(
			input *requestPresenter.AddCandidateRequest,
			output *responsePresenter.AddNewCandidateResponse,
		) (mvc.Result, error)
	}

	AddNewCandidateUseCase struct {
		AddNewCandidateService adminService.IAddNewCandidate
		ActionResultService    actionResultService.IActionResult
	}
)

func (this *AddNewCandidateUseCase) Execute(
	input *requestPresenter.AddCandidateRequest,
	output *responsePresenter.AddNewCandidateResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" || input.InputCandidate == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	inputCandidate := input.InputCandidate
	inputCandidate.SigningInfo = new(model.CandidateSigningInfo)
	inputCandidate.Version = libCommon.PointerPrimitive(time.Now())

	err := this.AddNewCandidateService.Execute(input.CampaignUUID, inputCandidate)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResultService.
		Prepare().
		SetCode(201).
		ServeResponse(output)
}
