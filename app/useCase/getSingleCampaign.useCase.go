package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	libCommon "app/lib/common"
	adminService "app/service/admin"
	"errors"
)

type (
	IGetSingleCampaign interface {
		Execute(
			*requestPresenter.GetSingleCampaignRequest,
			*responsePresenter.GetSingleCampaignResponse,
		) (*responsePresenter.GetSingleCampaignResponse, error)
	}

	GetSingleCampaignUseCase struct {
		GetSingleCampaignService adminService.IGetCampaign
	}
)

func (this *GetSingleCampaignUseCase) Execute(
	input *requestPresenter.GetSingleCampaignRequest,
	output *responsePresenter.GetSingleCampaignResponse,
) (*responsePresenter.GetSingleCampaignResponse, error) {

	if libCommon.Or(input == nil, input.UUID == "") {

		return nil, errors.New("invalid input")
	}

	data, err := this.GetSingleCampaignService.Execute(input.UUID)

	if err != nil {

		return nil, err
	}

	// response := responsePresenter.GetSingleCampaignResponse{
	// 	Message: "success",
	// 	Data:    *data,
	// }

	// resContent, err := json.Marshal(response)

	// if err != nil {

	// 	return nil, err
	// }

	output.Message = "succes"
	output.Data = data

	return output, nil
}
