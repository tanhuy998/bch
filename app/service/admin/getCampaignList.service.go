package adminService

import (
	"app/domain/model"
	"app/repository"
)

type (
	IGetCampaignList interface {
		Execute(page int) ([]*model.Campaign, error)
	}

	AdminGetCampaignListService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignListService) Execute(page int) ([]*model.Campaign, error) {

	return this.CampaignRepo.GetCampaignList(int64(page))
}
