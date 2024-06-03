package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	libCommon "app/lib/common"
	actionResultService "app/service/actionResult"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetSingleCampaign interface {
		Execute(
			*requestPresenter.GetSingleCampaignRequest,
			*responsePresenter.GetSingleCampaignResponse,
		) (mvc.Result, error)
	}

	GetSingleCampaignUseCase struct {
		GetSingleCampaignService adminService.IGetCampaign
		ActionResult             actionResultService.IActionResult
	}
)

func (this *GetSingleCampaignUseCase) Execute(
	input *requestPresenter.GetSingleCampaignRequest,
	output *responsePresenter.GetSingleCampaignResponse,
) (mvc.Result, error) {

	if libCommon.Or(input == nil, input.UUID == "") {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	data, err := this.GetSingleCampaignService.Serve(input.UUID)

	if err != nil {

		//return nil, err
		return this.ActionResult.ServeErrorResponse(err)
	}

	// response := responsePresenter.GetSingleCampaignResponse{
	// 	Message: "success",
	// 	Data:    *data,
	// }

	output.Message = "succes"
	output.Data = data

	// res := NewResponse()

	// err = MarshalResponseContent(output, res)

	// // resContent, err := json.Marshal(response)

	// if err != nil {

	// 	return nil, err
	// }

	//res.Code = 200
	// return res, nil

	return this.ActionResult.ServeResponse(output)
}
