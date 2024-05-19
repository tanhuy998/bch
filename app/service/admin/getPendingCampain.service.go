package adminService

import (
	"app/domain/model"
	"app/repository"
)

type (
	IGetPendingCampaigns interface {
		Execute(page int) ([]*model.Campaign, error)
	}

	AdminGetPendingCampaigns struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetPendingCampaigns) Execute(page int) ([]*model.Campaign, error) {

	return this.CampaignRepo.GetPendingCampaigns(page, nil)
}
