package controller

import (
	"app/app/model"
	adminService "app/app/service/admin"
	authService "app/app/service/auth"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type CampaignController struct {
	DeleteCampaignOperation    adminService.IDeleteCampaign
	LaunchNewCampaignOperation adminService.ILaunchNewCampaign
	ModifyExistingOperation    adminService.IModifyExistingCampaign
}

/*
GET /campaign/{uuid:string}?p={number}
*/
func (this *CampaignController) GetCampaign(auth authService.IAuthService) string {

	fmt.Println(auth)

	return fmt.Sprintf("type %s", reflect.TypeOf(auth).String())
}

func (this *CampaignController) GetCampaignListOnPage() {

}

func (this *CampaignController) NewCampaign(ctx iris.Context, inputCampaign *model.Campaign) mvc.Response {

	//repository.TestCampaignRepo()
	err := this.LaunchNewCampaignOperation.Execute(inputCampaign)

	if err != nil {

		return BadRequest(err)
	}

	return Created()
}

func (this *CampaignController) UpdateCampaign(uuid string, inputModel *model.Campaign) mvc.Response {

	err := this.ModifyExistingOperation.Execute(uuid, inputModel)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CampaignController) DeleteCampaign(uuid string) mvc.Response {

	err := this.DeleteCampaignOperation.Execute(uuid)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}
