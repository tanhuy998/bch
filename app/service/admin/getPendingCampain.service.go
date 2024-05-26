package adminService

import (
	"app/domain/model"
	"app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetPendingCampaigns interface {
		Serve(_id string, limit int, isPrevDir bool) (*repository.PaginationPack[model.Campaign], error)
	}

	AdminGetPendingCampaigns struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetPendingCampaigns) Serve(
	_id string, limit int, isPrevDir bool,
) (*repository.PaginationPack[model.Campaign], error) {

	objID, err := primitive.ObjectIDFromHex(_id)

	if err != nil {

		objID = primitive.NilObjectID
	}

	data, err := this.CampaignRepo.GetPendingCampaigns(objID, int64(limit), isPrevDir, nil)

	if err != nil {

		return nil, err
	}

	return data, nil
}
