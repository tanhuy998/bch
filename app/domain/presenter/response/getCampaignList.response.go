package responsePresenter

import (
	"app/domain/model"
)

type (
	GetCampaignListResponse struct {
		Message    string               `json:"message"`
		Data       []*model.Campaign    `json:"data"`
		Navigation PaginationNavigation `json:"navigation"`
	}
)

func (this *GetCampaignListResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}
