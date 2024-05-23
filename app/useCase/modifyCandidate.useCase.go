package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IModifyCandidate interface {
		Execute(
			input *requestPresenter.ModifyCandidateRequest,
			output *responsePresenter.ModifyCandidateResponse,
		) (mvc.Result, error)
	}

	ModifyCandidateUseCase struct {
		ModifyCandidateService adminService.IModifyCandidate
	}
)

func (this *ModifyCandidateUseCase) Execute(
	input *requestPresenter.ModifyCandidateRequest,
	output *responsePresenter.ModifyCandidateResponse,
) (mvc.Result, error) {

	input.Candidate.UUID = nil

	err := this.ModifyCandidateService.Execute(input.UUID, input.Candidate)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"

	MarshalResponseContent(output, res)

	return res, nil
}
