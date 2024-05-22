package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IUpdateCampaign interface {
		Execute(*requestPresenter.UpdateCampaignRequest, *responsePresenter.UpdateCampaignResponse) (mvc.Result, error)
	}

	UpdateCampaignUseCase struct {
		ModifyCampaignService adminService.IModifyExistingCampaign
	}
)

func (this *UpdateCampaignUseCase) Execute(
	input *requestPresenter.UpdateCampaignRequest,
	output *responsePresenter.UpdateCampaignResponse,
) (mvc.Result, error) {

	var (
		uuid       string          = input.UUID
		inputModel *model.Campaign = input.Data
	)

	if libCommon.Or(uuid == "", inputModel == nil) {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	err := this.ModifyCampaignService.Execute(uuid, inputModel)

	if err != nil {

		return nil, err
	}

	res := newResponse()
	output.Message = "success"
	err = marshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	res.Code = 200

	return res, nil
}
