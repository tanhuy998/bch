package adminService

import (
	"app/domain/model"
	"app/internal/common"
	"app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetCampaignList interface {
		Execute(_id string, limit int, isPrevDir bool) (data []*model.Campaign, pageNumber common.PaginationPage, err error)
	}

	AdminGetCampaignListService struct {
		CampaignRepo repository.ICampaignRepository
	}
)

func (this *AdminGetCampaignListService) Execute(
	_id string, limit int, isPrevDir bool,
) (data []*model.Campaign, pageNumber common.PaginationPage, err error) {

	objID, err := primitive.ObjectIDFromHex(_id)

	if err != nil {

		objID = primitive.NilObjectID
	}

	data, docCount, err := this.CampaignRepo.GetCampaignList(objID, int64(limit), isPrevDir, nil)

	if err != nil {

		return nil, 0, err
	}

	return data, calculatePageNumber(objID, docCount), nil
}

func calculatePageNumber(_id primitive.ObjectID, docCount int64) common.PaginationPage {

	return common.PAGINATION_FIRST_PAGE
}
