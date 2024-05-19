package adminService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
)

type (
	IGetCampaign interface {
		Execute(uuid string) (*model.Campaign, error)
	}

	AdminGetCampaignService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignService) Execute(string) (*model.Campaign, error) {

	uuid := uuid.New()

	return this.CampaignRepo.FindByUUID(uuid, nil)
}
