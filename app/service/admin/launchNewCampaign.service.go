package adminService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
)

type (
	ILaunchNewCampaign interface {
		Execute(*model.Campaign) error
	}

	AdminLaunchNewCampaignService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminLaunchNewCampaignService) Execute(model *model.Campaign) error {

	return this.launchNewCampaign(model)
}

func (this *AdminLaunchNewCampaignService) launchNewCampaign(model *model.Campaign) error {

	model.UUID = uuid.New()

	return this.CampaignRepo.Create(model, nil)
}
