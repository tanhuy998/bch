package adminService

import (
	"app/src/model"
	"app/src/repository"

	"github.com/google/uuid"
)

type (
	IGetCampaign interface {
		Serve(uuid string) (*model.Campaign, error)
	}

	AdminGetCampaignService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignService) Serve(inputUUID string) (*model.Campaign, error) {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return nil, err
	}

	return this.CampaignRepo.FindByUUID(uuid, nil)
}
