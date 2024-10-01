package usecase

import (
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

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
		ActionResultService       actionResultService.IActionResult
	}
)

func (this *GetSingleCandidateByUUIDUseCase) Execute(
	input *requestPresenter.GetSingleCandidateRequest,
	output *responsePresenter.GetSingleCandidateResponse,
) (mvc.Result, error) {

	if input.UUID == "" {

		return this.ActionResultService.ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	candidate, err := this.GetSingleCandidateService.Serve(input.UUID)

	if err != nil {

		return this.ActionResultService.ServeErrorResponse(err)
	}

	if candidate == nil {

		return this.ActionResultService.ServeErrorResponse(common.ERR_HTTP_NOT_FOUND)
	}

	output.Data = candidate
	output.Message = "success"

	return this.ActionResultService.ServeResponse(output)
}
