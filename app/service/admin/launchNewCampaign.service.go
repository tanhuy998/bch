package adminService

import (
	"app/domain/model"
	"app/repository"
	"errors"

	"github.com/google/uuid"
)

type (
	ILaunchNewCampaign interface {
		Execute(*model.Campaign) (*uuid.UUID, error)
	}

	AdminLaunchNewCampaignService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminLaunchNewCampaignService) Execute(model *model.Campaign) (*uuid.UUID, error) {

	uuid, err := this.launchNewCampaign(model)

	if err != nil {

		return nil, err
	}

	return uuid, err
}

func (this *AdminLaunchNewCampaignService) launchNewCampaign(model *model.Campaign) (*uuid.UUID, error) {

	if !model.Expire.After(*model.IssueTime) {

		return nil, errors.New("campaing expire time be a day in the future of issue time")
	}

	newUUID := uuid.New()
	model.UUID = &newUUID

	err := this.CampaignRepo.Create(model, nil)

	if err != nil {

		return nil, err
	}

	return &newUUID, err
}
