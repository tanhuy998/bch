package usecase

import (
	adminService "app/service/admin"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IDeleteCampaign interface {
		Execute(
			*requestPresenter.DeleteCampaignRequest,
			*responsePresenter.DeleteCampaignResponse,
		) (mvc.Result, error)
	}

	DeleteCampaignUseCase struct {
		DeleteCampaignService adminService.IDeleteCampaign
	}
)

func (this *DeleteCampaignUseCase) Execute(
	input *requestPresenter.DeleteCampaignRequest,
	output *responsePresenter.DeleteCampaignResponse,
) (mvc.Result, error) {

	var (
		uuid string = input.UUID
		err  error
	)

	err = this.DeleteCampaignService.Execute(uuid)

	if err != nil {

		return nil, err
	}

	output.Message = "success"
	res := NewResponse()

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	res.Code = 200

	return res, nil
}
