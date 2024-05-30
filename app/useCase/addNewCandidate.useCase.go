package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	adminService "app/service/admin"
	"time"

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

	if input.InputCandidate == nil {

		return nil, common.ERR_BAD_REQUEST
	}

	inputCandidate := input.InputCandidate
	inputCandidate.SigningInfo = new(model.CandidateSigningInfo)
	inputCandidate.Version = libCommon.PointerPrimitive(time.Now())

	err := this.AddNewCandidateService.Execute(input.CampaignUUID, inputCandidate)

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
