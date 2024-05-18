package adminService

import (
	"app/app/model"
	"app/app/repository"

	"github.com/google/uuid"
)

type (
	IModifyExistingCampaign interface {
		Execute(string, *model.Campaign) error
	}

	AdminModifyExistingCampaign struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminModifyExistingCampaign) Execute(inputUUID string, model *model.Campaign) error {

	uuid, err := uuid.Parse(inputUUID)

	if err != nil {

		return err
	}

	return this.modifyExistingCampaign(uuid, model)
}

func (this *AdminModifyExistingCampaign) modifyExistingCampaign(uuid uuid.UUID, model *model.Campaign) error {

	model.UUID = uuid

	return this.CampaignRepo.Update(model, nil)
}
