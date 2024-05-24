package adminService

import (
	"app/domain/model"
	"app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetCampaignList interface {
		Execute(_id string, limit int, direction int) ([]*model.Campaign, error)
	}

	AdminGetCampaignListService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignListService) Execute(_id string, limit int, direction int) ([]*model.Campaign, error) {

	objID, err := primitive.ObjectIDFromHex(_id)

	if err != nil {

		return nil, err
	}

	return this.CampaignRepo.GetCampaignList(objID, int64(limit), int64(direction))
}
