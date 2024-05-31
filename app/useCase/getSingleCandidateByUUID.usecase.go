package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCandidateByUUID interface {
		Execute(
			input *requestPresenter.GetSingleCandidateRequest,
			output *responsePresenter.GetSingleCandidateResponse,
		) (mvc.Result, error)
	}

	GetSingleCandidateByUUIDUseCase struct {
		GetSingleCandidateService adminService.IGetSingleCandidateByUUID
	}
)

func (this *GetSingleCandidateByUUIDUseCase) Execute(
	input *requestPresenter.GetSingleCandidateRequest,
	output *responsePresenter.GetSingleCandidateResponse,
) (mvc.Result, error) {

	if input.UUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	candidate, err := this.GetSingleCandidateService.Serve(input.UUID)

	if err != nil {

		return nil, err
	}

	if candidate == nil {

		return nil, common.ERR_HTTP_NOT_FOUND
	}

	res := NewResponse()
	output.Data = candidate
	output.Message = "success"

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
