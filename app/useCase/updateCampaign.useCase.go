package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	adminService "app/service/admin"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IUpdateCampaign interface {
		Execute(*requestPresenter.UpdateCampaignRequest, *responsePresenter.UpdateCampaignResponse) (mvc.Result, error)
	}

	UpdateCampaignUseCase struct {
		GetSingleCampaignService adminService.IGetCampaign
		ModifyCampaignService    adminService.IModifyExistingCampaign
	}
)

func (this *UpdateCampaignUseCase) Execute(
	input *requestPresenter.UpdateCampaignRequest,
	output *responsePresenter.UpdateCampaignResponse,
) (mvc.Result, error) {

	var (
		uuid            string          = input.UUID
		campaignUpdated *model.Campaign = new(model.Campaign)
	)

	inputModel := input.Data

	if libCommon.Or(uuid == "", inputModel == nil) {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	targetCampaingn, err := this.GetSingleCampaignService.Serve(uuid)

	if err != nil || targetCampaingn == nil {

		return nil, errors.New("not found")
	}

	if inputModel.Title != nil || *inputModel.Title != "" {

		campaignUpdated.Title = inputModel.Title

	} else {

		campaignUpdated.Title = nil
	}

	campaignUpdated.Expire = inputModel.Expire

	err = this.ModifyCampaignService.Serve(uuid, campaignUpdated)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"
	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	res.Code = 200

	return res, nil
}
