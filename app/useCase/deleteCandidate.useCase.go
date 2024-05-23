package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IDeleteCandidate interface {
		Execute(
			input *requestPresenter.DeleteCandidateRequest,
			output *responsePresenter.DeleteCandidateResponse,
		) (mvc.Result, error)
	}

	DeleteCandidateUseCase struct {
		DeleteCandidateService adminService.IDeleteCandidate
	}
)

func (this *DeleteCandidateUseCase) Execute(
	input *requestPresenter.DeleteCandidateRequest,
	output *responsePresenter.DeleteCandidateResponse,
) (mvc.Result, error) {

	if input.CandidateUUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	err := this.DeleteCandidateService.Execute(input.CandidateUUID)

	if err != nil {

		return nil, err
	}

	res := newResponse()
	output.Message = "success"

	err = marshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
