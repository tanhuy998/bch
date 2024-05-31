package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
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
	}
)

func (this *ModifyExistingCandidateUseCase) Execute(
	input *requestPresenter.ModifyExistingCandidateRequest,
	output *responsePresenter.ModifyExistingCandidateResponse,
) (mvc.Result, error) {

	if input.UUID == "" {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	inputData := input.Data

	var updatedCandidate *model.Candidate = &model.Candidate{
		Name:     inputData.Name,
		IDNumber: inputData.IDNumber,
		Phone:    inputData.Phone,
		Address:  inputData.Address,
	}

	err := this.ModifyCandidateService.Serve(input.UUID, updatedCandidate)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"

	MarshalResponseContent(output, res)

	return res, nil
}
