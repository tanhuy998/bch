package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	actionResultService "app/service/actionResult"
	"app/service/signingService"

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
		//GetSingleCanidateSigingInfoServoce candidateService.IGetSingleCandidateSigningInfo
		ActionResultService           actionResultService.IActionResult
		GetSingleCandidateSigningInfo signingService.IGetSingleCandidateSigningInfo
	}
)

func (this *GetSingleCandidateSigningInfoUseCase) Execute(
	input *requestPresenter.GetSingleCandidateSigningInfoRequest,
	output *responsePresenter.GetSingleCandidateSigningInfoResponse,
) (mvc.Result, error) {

	if input.CandidateUUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	// singingInfo, err := this.GetSingleCanidateSigingInfoServoce.Serve(input.CampaignUUID, input.CandidateUUID)
	singingInfo, err := this.GetSingleCandidateSigningInfo.Serve(input.CandidateUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	if singingInfo == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_HTTP_NOT_FOUND)
	}

	output.Message = "success"
	output.Data = singingInfo

	return this.ActionResultService.ServeResponse(output)
}

func HideSensitiveInfo(signingInfo *model.CandidateSigningInfo) {

	signingInfo.CivilIndentity.IDNumber = ""
}
