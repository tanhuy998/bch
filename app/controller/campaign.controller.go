package controller

import (
	"app/app/config"
	libCommon "app/app/lib/common"
	"app/app/model"
	"app/app/repository"
	authService "app/app/service/auth"
	"fmt"
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
)

type CampaignController struct {
}

/*
GET /campaign/{uuid:string}?p={number}
*/
func (this *CampaignController) GetCampaign(auth authService.IAuthService) string {

	fmt.Println(auth)

	return fmt.Sprintf("type %T", reflect.TypeOf(auth))
}

func (this *CampaignController) GetCampaignListOnPage() {

}

func (this *CampaignController) NewCampaign(ctx iris.Context, campaignRepo repository.ICampaignRepository) {

	//repository.TestCampaignRepo()

	reqBody, ok := ctx.Values().Get(config.REQUEST_BODY).(*model.Campaign)

	var newCampaign *model.Campaign = libCommon.Ternary(ok, reqBody, new(model.Campaign))

	newCampaign.IssueTime = libCommon.PointerPrimitive(time.Now())

	campaignRepo.Create(newCampaign)
}

func (this *CampaignController) UpdateCampaign() {

}

func (this *CampaignController) DeleteCampaign() {

}

func (this *CandidateController) HealthCheck() {

}

func (this *CandidateController) H() {

}
