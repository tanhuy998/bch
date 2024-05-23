package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/port"
	usecase "app/useCase"
	"fmt"

	"github.com/kataras/iris/v12/mvc"
)

const (
	CONTENT_TYPE = "application/json"
)

type CampaignController struct {
	ErrorResponseUseCase          port.IActionResult
	GetPendingCampaignsUseCase    usecase.IGetPendingCampaigns
	GetCampaignListUseCase        usecase.IGetCampaignList
	GetSingleCampaignUseCase      usecase.IGetSingleCampaign
	DeleteCampaignUseCase         usecase.DeleteCampaignUseCase
	LaunNewCampaignUseCase        usecase.ILaunchNewCampaign
	UpdateExistingCampaignUseCase usecase.IUpdateCampaign
}

func (this *CampaignController) HandleHTTPError(err mvc.Err, statusCode mvc.Code) *mvc.Response {

	var msg string

	if err != nil {

		msg = err.Error()
	} else {

		msg = "error"
	}

	res := this.ErrorResponseUseCase.NewActionResponse()
	errOutput := &(responsePresenter.ErrorResponse{
		Message: msg,
	})

	this.ErrorResponseUseCase.MarshallOutput(errOutput, res)

	return res
}

/*
GET /campaign/{uuid:string}?p={number}
*/
func (this *CampaignController) GetCampaign(
	input *requestPresenter.GetSingleCampaignRequest,
	output *responsePresenter.GetSingleCampaignResponse,
) (mvc.Result, error) {

	fmt.Printf("c %s \n", input.UUID)
	return this.GetSingleCampaignUseCase.Execute(input, output)
}

func (this *CampaignController) GetCampaignListOnPage(
	input *requestPresenter.GetCampaignListRequest,
	output *responsePresenter.GetCampaignListResponse,
) (mvc.Result, error) {

	return this.GetCampaignListUseCase.Execute(input, output)
}

func (this *CampaignController) GetPendingCampaigns(
	input *requestPresenter.GetPendingCampaignRequest,
	output *responsePresenter.GetPendingCampaingsResponse,
) (mvc.Result, error) {

	return this.GetPendingCampaignsUseCase.Execute(input, output)
}

func (this *CampaignController) NewCampaign(
	input *requestPresenter.LaunchNewCampaignRequest,
	output *responsePresenter.LaunchNewCampaignResponse,
) (mvc.Result, error) {

	//repository.TestCampaignRepo()

	return this.LaunNewCampaignUseCase.Execute(input, output)
}

func (this *CampaignController) UpdateCampaign(
	input *requestPresenter.UpdateCampaignRequest,
	output *responsePresenter.UpdateCampaignResponse,
) (mvc.Result, error) {

	return this.UpdateExistingCampaignUseCase.Execute(input, output)
}

func (this *CampaignController) DeleteCampaign(
	input *requestPresenter.DeleteCampaignRequest,
	output *responsePresenter.DeleteCampaignResponse,
) (mvc.Result, error) {

	return this.DeleteCampaignUseCase.Execute(input, output)
}
