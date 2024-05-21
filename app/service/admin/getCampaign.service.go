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

func (this *AdminGetCampaignService) Execute(inputUUID string) (*model.Campaign, error) {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return nil, err
	}

	return this.CampaignRepo.FindByUUID(uuid, nil)
}
