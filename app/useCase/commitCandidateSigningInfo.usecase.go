package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"
	candidateService "app/service/candidate"
	"fmt"

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
		GetSingleCandidateService         adminService.IGetSingleCandidateByUUID
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
	fmt.Println("a")
	_, _, err := this.Retrieve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return nil, err
	}
	fmt.Println("b")
	err = this.CommitCandidateSigningInfoService.Serve(input.CandidateUUID, input.CampaignUUID, input.Data)

	if err != nil {

		return nil, err
	}
	fmt.Println("c")
	res := NewResponse()

	updated, err := this.GetSingleCandidateByUUIDService.Serve(input.CampaignUUID)

	if err != nil {

		fmt.Println(err)
	}

	output.UpdatedData = updated

	output.Message = "success"

	err = MarshalResponseContent(output, res)
	fmt.Println("d")
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
	fmt.Println(1)
	candidate, err := this.GetSingleCandidateByUUIDService.Serve(candidateUUID_str)

	if err != nil {

		return nil, nil, err
	}
	fmt.Println(2)
	return campaign, candidate, nil
}
