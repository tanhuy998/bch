package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

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
	}
)

func (this *AddNewCandidateUseCase) Execute(
	input *requestPresenter.AddCandidateRequest,
	output *responsePresenter.AddNewCandidateResponse,
) (mvc.Result, error) {

	if input.CampaignUUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	err := this.AddNewCandidateService.Execute(input.CampaignUUID, input.CandidateDetail)

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
