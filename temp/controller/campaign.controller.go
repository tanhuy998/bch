package controller

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"

	"app/src/infrastructure/http/middleware"
	usecase "app/useCase"
	"fmt"

	"github.com/kataras/iris/v12/mvc"
)

const (
	CONTENT_TYPE = "application/json"
)

type CampaignController struct {
	// ErrorResponseUseCase          port.IActionResult
	GetPendingCampaignsUseCase    usecase.IGetPendingCampaigns
	GetCampaignListUseCase        usecase.IGetCampaignList
	GetSingleCampaignUseCase      usecase.IGetSingleCampaign
	DeleteCampaignUseCase         usecase.DeleteCampaignUseCase
	LaunNewCampaignUseCase        usecase.ILaunchNewCampaign
	UpdateExistingCampaignUseCase usecase.IUpdateCampaign
	CampaignProgresss             usecase.ICampaignProgress
}

func (this *CampaignController) BeforeActivation(activator mvc.BeforeActivation) {

	container := activator.Router().ConfigureContainer().Container

	activator.Handle(
		"GET", "/", "GetCampaignListOnPage",

		middleware.BindPresenters[requestPresenter.GetCampaignListRequest, responsePresenter.GetCampaignListResponse](container),
	).SetName("GET_LIST_CAMPAIGNS")

	activator.Handle(
		"GET", "/{uuid:uuid}", "GetCampaign",

		middleware.BindPresenters[requestPresenter.GetSingleCampaignRequest, responsePresenter.GetSingleCampaignResponse](container),
	).SetName("GET_SINGLE_CAMPAIGN")

	activator.Handle(
		"GET", "/{uuid:uuid}/progress", "GetCampaignProgress",
		middleware.BindPresenters[requestPresenter.CampaignProgressRequestPresenter, responsePresenter.CampaignProgressResponsePresenter](container),
	)

	activator.Handle(
		"GET", "/pending", "GetPendingCampaigns",

		middleware.BindPresenters[requestPresenter.GetPendingCampaignRequest, responsePresenter.GetPendingCampaingsResponse](container),
	).SetName("GET_PENDING_CAMPAIGNS")

	activator.Handle(
		"POST", "/", "NewCampaign",

		middleware.BindPresenters[requestPresenter.LaunchNewCampaignRequest, responsePresenter.LaunchNewCampaignResponse](container),
	).SetName("LAUNCH_NEW_CAMPAIGN")

	activator.Handle(
		"PATCH", "/{uuid:uuid}", "UpdateCampaign",

		middleware.BindPresenters[requestPresenter.UpdateCampaignRequest, responsePresenter.UpdateCampaignResponse](container),
	).SetName("UPDATE_CAMPAIGN")

	// activator.Handle(
	// 	"DELETE", "/{uuid:uuid}", "DeleteCampaign",
	// 	middleware.Authorize(authService.AuthorizationLicense{
	// 		Fields: campaignField,
	// 		Claims: []authService.AuthorizationClaim{auth_post_claim},
	// 		//Groups: []authService.AuthorizationGroup{auth_commander_group},
	// 	}),
	// 	middleware.BindPresenters[requestPresenter.DeleteCampaignRequest, responsePresenter.DeleteCampaignResponse](container),
	// ).SetName("DELETE_CAMPAIGN")

	activator.Handle(
		"PATCH", "/test/{uuid:uuid}", "TestPatch",

		middleware.BindPresenters[requestPresenter.UpdateCampaignRequest, responsePresenter.UpdateCampaignResponse](container),
	)
}

// func (this *CampaignController) HandleHTTPError(err mvc.Err, statusCode mvc.Code) *mvc.Response {

// 	var msg string

// 	if err != nil {

// 		msg = err.Error()
// 	} else {

// 		msg = "error"
// 	}

// 	res := this.ErrorResponseUseCase.NewActionResponse()
// 	errOutput := &(responsePresenter.ErrorResponse{
// 		Message: msg,
// 	})

// 	res.Code = 400
// 	this.ErrorResponseUseCase.MarshallOutput(errOutput, res)

// 	return res
// }

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

func (this *CampaignController) TestPatch(
	input *requestPresenter.UpdateCampaignRequest,
) string {

	//fmt.Println(input)
	return "true"
}

func (this *CampaignController) GetCampaignListOnPage(
	input *requestPresenter.GetCampaignListRequest,
	output *responsePresenter.GetCampaignListResponse,
) (mvc.Result, error) {
	fmt.Println("get campaign list controller")
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

func (this *CampaignController) GetCampaignProgress(
	input *requestPresenter.CampaignProgressRequestPresenter,
	output *responsePresenter.CampaignProgressResponsePresenter,
) (mvc.Result, error) {

	return this.CampaignProgresss.Execute(input, output)
}
