package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	candidateService "app/service/candidate"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICommitCandidateSigningInfo interface {
		Execute(
			input *requestPresenter.CommitCandidateSigningInfoRequest,
			outout *responsePresenter.CommitCandidateSigningInfoResponse,
		) (mvc.Result, error)
	}

	CommitCandidateSigningInfoUseCase struct {
		CommitCandidateSigningInfoService candidateService.ICommitCandidateSigningInfo
		CommitLoggerService               candidateService.ICandidateSigningCommitLogger
		CheckSigningExistence             candidateService.ICheckSigningExistence
		GetSingleCandidateByUUIDService   adminService.IGetSingleCandidateByUUID
		GetSingleCandidateService         adminService.IGetSingleCandidateByUUID
		GetCampaignService                adminService.IGetCampaign
		ActionResultService               actionResultService.IActionResult
	}
)

func (this *CommitCandidateSigningInfoUseCase) Execute(
	input *requestPresenter.CommitCandidateSigningInfoRequest,
	output *responsePresenter.CommitCandidateSigningInfoResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" || input.CandidateUUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	_, candidate, err := this.Retrieve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	err = this.WriteCommitLog(input.Data, candidate)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	err = this.CommitCandidateSigningInfoService.Serve(input.CandidateUUID, input.CampaignUUID, input.Data)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	updated, err := this.GetSingleCandidateByUUIDService.Serve(input.CampaignUUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	output.UpdatedData = updated
	output.Message = "success"

	return this.ActionResultService.ServeResponse(output)
}

func (this *CommitCandidateSigningInfoUseCase) WriteCommitLog(
	inputCommit *model.CandidateSigningInfo,
	candidate *model.Candidate,
) error {

	return this.CommitLoggerService.CompareAndServe(inputCommit, candidate)
}

func (this *CommitCandidateSigningInfoUseCase) Retrieve(
	campaignUUID_str string,
	candidateUUID_str string,
) (*model.Campaign, *model.Candidate, error) {

	campaign, err := this.GetCampaignService.Serve(campaignUUID_str)

	if err != nil {

		return nil, nil, err
	}

	candidate, err := this.GetSingleCandidateByUUIDService.Serve(candidateUUID_str)

	if err != nil {

		return nil, nil, err
	}

	return campaign, candidate, nil
}
