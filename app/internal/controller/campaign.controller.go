package controller

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	libCommon "app/lib/common"
	adminService "app/service/admin"
	"encoding/json"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

const (
	CONTENT_TYPE = "application/json"
)

type CampaignController struct {
	GetPendingCampaignOperation adminService.IGetPendingCampaigns
	GetCampaignListOperation    adminService.IGetCampaignList
	GetCampaignOperation        adminService.IGetCampaign
	DeleteCampaignOperation     adminService.IDeleteCampaign
	LaunchNewCampaignOperation  adminService.ILaunchNewCampaign
	ModifyExistingOperation     adminService.IModifyExistingCampaign
}

/*
GET /campaign/{uuid:string}?p={number}
*/
func (this *CampaignController) GetCampaign(presenter *requestPresenter.GetSingleCampaignRequest) mvc.Response {

	if libCommon.Or(presenter == nil, presenter.UUID == "") {

		return BadRequest(errors.New("invalid input"))
	}

	data, err := this.GetCampaignOperation.Execute(presenter.UUID)

	if err != nil {

		return BadRequest(err)
	}

	response := responsePresenter.GetSingleCampaignResponse{
		Message: "success",
		Data:    *data,
	}

	resContent, err := json.Marshal(response)

	if err != nil {

		return BadRequest(err)
	}

	return mvc.Response{
		Code:        200,
		ContentType: "application",
		Content:     resContent,
	}
}

func (this *CampaignController) GetCampaignListOnPage(presenter *requestPresenter.GetCampaignListRequest) mvc.Response {

	if presenter == nil {

		return BadRequest(errors.New("Invalid input"))
	}

	_, err := this.GetCampaignListOperation.Execute(presenter.PageNumber)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CampaignController) GetPendingCampaigns(presenter *requestPresenter.GetPendingCampaignRequest) mvc.Response {

	if presenter == nil {

		return BadRequest(errors.New("Invalid input"))
	}

	_, err := this.GetPendingCampaignOperation.Execute(presenter.PageNumber)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CampaignController) NewCampaign(presenter *requestPresenter.LaunchNewCampaignRequest) mvc.Response {

	//repository.TestCampaignRepo()
	var inputCampaign *model.Campaign = presenter.Data

	if inputCampaign == nil {

		return BadRequest(errors.New("invalid input"))
	}

	err := this.LaunchNewCampaignOperation.Execute(inputCampaign)

	if err != nil {

		return BadRequest(err)
	}

	return Created()
}

func (this *CampaignController) UpdateCampaign(presenter *requestPresenter.UpdateCampaignRequest) mvc.Response {

	var (
		uuid       string          = presenter.UUID
		inputModel *model.Campaign = presenter.Data
	)

	if libCommon.Or(uuid == "", inputModel == nil) {

		return BadRequest(errors.New("invalid input"))
	}

	err := this.ModifyExistingOperation.Execute(uuid, inputModel)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CampaignController) DeleteCampaign(presenter *requestPresenter.DeleteCampaignRequest) mvc.Response {

	var (
		uuid string = presenter.UUID
	)

	err := this.DeleteCampaignOperation.Execute(uuid)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}
