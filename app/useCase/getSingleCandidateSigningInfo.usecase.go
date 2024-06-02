package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	candidateService "app/service/candidate"
	"fmt"

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
	}
)

func (this *GetSingleCandidateSigningInfoUseCase) Execute(
	input *requestPresenter.GetSingleCandidateSigningInfoRequest,
	output *responsePresenter.GetSingleCandidateSigningInfoResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" || input.CandidateUUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	singingInfo, err := this.GetSingleCanidateSigingInfoServoce.Serve(input.CampaignUUID, input.CandidateUUID)

	if err != nil {

		return nil, err
	}

	res := NewResponse()

	output.Message = "success"
	output.Data = singingInfo
	fmt.Println(output)
	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
