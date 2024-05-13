package controller

import (
	authService "app/app/service/auth"
	"fmt"
	"reflect"
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

func (this *CampaignController) NewCampaign() {

}

func (this *CampaignController) UpdateCampaign() {

}

func (this *CampaignController) DeleteCampaign() {

}

func (this *CandidateController) HealthCheck() {

}

func (this *CandidateController) H() {

}
