package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	"app/repository"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	RetrievedData_T = repository.PaginationPack[model.Campaign]

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

	dataPack, err := this.GetCampaignListService.Serve(input.PivotID, input.PageSizeLimit, input.IsPrev)

	if err != nil {

		return nil, err
	}

	pageNumber := common.PaginationPage(1)

	res := NewResponse()
	output.Message = "success"
	output.Data = dataPack.Data

	resolveNext(output, dataPack, pageNumber)
	resolvePrev(output, dataPack, pageNumber)

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}

func resolveNext(
	output *responsePresenter.GetCampaignListResponse,
	dataPack *RetrievedData_T,
	pageNumber common.PaginationPage,
) {

	lastIndex := len(dataPack.Data) - 1

	if lastIndex <= 0 {

		return
	}

	output.Navigation.Next = dataPack.Data[lastIndex].ObjectID.Hex()
}

func resolvePrev(
	output *responsePresenter.GetCampaignListResponse,
	retrievedData *RetrievedData_T,
	pageNumber common.PaginationPage,
) {

}
