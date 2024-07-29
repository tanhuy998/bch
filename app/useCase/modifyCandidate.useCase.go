package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IModifyExistingCandidate interface {
		Execute(
			input *requestPresenter.ModifyExistingCandidateRequest,
			output *responsePresenter.ModifyExistingCandidateResponse,
		) (mvc.Result, error)
	}

	ModifyExistingCandidateUseCase struct {
		ModifyCandidateService adminService.IModifyExistingCandidate
		ActionResultService    actionResultService.IActionResult
	}
)

func (this *ModifyExistingCandidateUseCase) Execute(
	input *requestPresenter.ModifyExistingCandidateRequest,
	output *responsePresenter.ModifyExistingCandidateResponse,
) (mvc.Result, error) {

	if input.UUID == "" {

		return this.ActionResultService.Prepare().ServeErrorResponse(common.ERR_INVALID_HTTP_INPUT)
	}

	inputData := input.Data

	// var updatedCandidate *model.Candidate = &model.Candidate{
	// 	Name:        inputData.Name,
	// 	IDNumber:    inputData.IDNumber,
	// 	Phone:       inputData.Phone,
	// 	Address:     inputData.Address,
	// 	DateOfBirth: inputData.DateOfBirth,
	// }

	// err := this.ModifyCandidateService.Serve(input.UUID, updatedCandidate)

	err := this.ModifyCandidateService.Serve(input.UUID, inputData)

	if err != nil {

		return this.ActionResultService.Prepare().ServeErrorResponse(err)
	}

	output.Message = "success"

	return this.ActionResultService.Prepare().ServeResponse(output)
}
