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
	RetrievedData_T []*model.Campaign

	IGetCampaignList interface {
		Execute(
			input *requestPresenter.GetCampaignListRequest,
			output *responsePresenter.GetCampaignListResponse,
		) (mvc.Result, error)
	}

	GetCampaignListUseCase struct {
		GetCampaignListService adminService.IGetCampaignList
	}
)

func (this *GetCampaignListUseCase) Execute(
	input *requestPresenter.GetCampaignListRequest,
	output *responsePresenter.GetCampaignListResponse,
) (mvc.Result, error) {

	if input == nil {

		return nil, common.ERR_INVALID_HTTP_INPUT
	}

	data, pageNumber, err := this.GetCampaignListService.Execute(input.PivotID, input.PageSizeLimit, input.IsPrev)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"
	output.Data = data

	resolveNext(output, data, pageNumber)
	resolvePrev(output, data, pageNumber)

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}

func resolveNext(
	output *responsePresenter.GetCampaignListResponse,
	retrievedData RetrievedData_T,
	pageNumber common.PaginationPage,
) {

	lastIndex := len(retrievedData) - 1

	output.Navigation.Next = retrievedData[lastIndex].ObjectID.Hex()
}

func resolvePrev(
	output *responsePresenter.GetCampaignListResponse,
	retrievedData RetrievedData_T,
	pageNumber common.PaginationPage,
) {

}
