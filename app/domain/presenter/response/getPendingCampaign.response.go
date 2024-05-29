package responsePresenter

import "app/domain/model"

type GetPendingCampaingsResponse struct {
	Message    string               `json:"message"`
	Data       []*model.Campaign    `json:"data"`
	Navigation PaginationNavigation `json:"navigation"`
}

func (this *GetPendingCampaingsResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}
