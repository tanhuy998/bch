package adminService

import (
	"app/domain/model"
	"app/internal/common"
	"app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetCampaignList interface {
		Serve(_id string, limit int, isPrevDir bool) (data *repository.PaginationPack[model.Campaign], err error)
	}

	AdminGetCampaignListService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignListService) Serve(
	_id string, limit int, isPrevDir bool,
) (data *repository.PaginationPack[model.Campaign], err error) {

	objID, err := primitive.ObjectIDFromHex(_id)

	if err != nil {

		objID = primitive.NilObjectID
	}

	data, err = this.CampaignRepo.GetCampaignList(objID, int64(limit), isPrevDir, nil)

	if err != nil {

		return nil, err
	}

	return data, nil
}

func calculatePageNumber(_id primitive.ObjectID, docCount int64) common.PaginationPage {

	return common.PAGINATION_FIRST_PAGE
}
