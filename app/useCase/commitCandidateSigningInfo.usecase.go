package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
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
		GetSingleCandidateByUUIDService   adminService.IGetSingleCandidateByUUID
		GetCampaignService                adminService.IGetCampaign
	}
)

func (this *CommitCandidateSigningInfoUseCase) Execute(
	input *requestPresenter.CommitCandidateSigningInfoRequest,
	output *responsePresenter.CommitCandidateSigningInfoResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" || input.CandidateUUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	_, _, err := this.Retrieve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return nil, err
	}

	err = this.CommitCandidateSigningInfoService.Serve(input.CandidateUUID, input.Data)

	if err != nil {

		return nil, err
	}

	res := NewResponse()

	output.Message = "success"

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}

func (this *CommitCandidateSigningInfoUseCase) Retrieve(campaignUUID_str string, candidateUUID_str string) (*model.Campaign, *model.Candidate, error) {

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
