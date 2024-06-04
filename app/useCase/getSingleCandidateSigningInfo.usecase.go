package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	actionResultService "app/service/actionResult"
	candidateService "app/service/candidate"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCandidateSigningInfo interface {
		Execute(
			input *requestPresenter.GetSingleCandidateSigningInfoRequest,
			output *responsePresenter.GetSingleCandidateSigningInfoResponse,
		) (mvc.Result, error)
	}

	GetSingleCandidateSigningInfoUseCase struct {
		GetSingleCanidateSigingInfoServoce candidateService.IGetSingleCandidateSigningInfo
		ActionResultService                actionResultService.IActionResult
	}
)

func (this *GetSingleCandidateSigningInfoUseCase) Execute(
	input *requestPresenter.GetSingleCandidateSigningInfoRequest,
	output *responsePresenter.GetSingleCandidateSigningInfoResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" || input.CandidateUUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	singingInfo, err := this.GetSingleCanidateSigingInfoServoce.Serve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	HideSensitiveInfo(singingInfo)

	output.Message = "success"
	output.Data = singingInfo

	return this.ActionResultService.ServeResponse(output)
}

func HideSensitiveInfo(signingInfo *model.CandidateSigningInfo) {

	signingInfo.CivilIndentity.IDNumber = ""
}
