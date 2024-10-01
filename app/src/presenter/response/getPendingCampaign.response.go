package responsePresenter

import "app/src/model"

type GetPendingCampaingsResponse struct {
	Message        string               `json:"message"`
	Data           []*model.Campaign    `json:"data"`
	DataTotalCount int64                `json:"dataTotalCount"`
	Navigation     PaginationNavigation `json:"navigation"`
}

func (this *GetPendingCampaingsResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}

func (this *GetPendingCampaingsResponse) SetTotalCount(count int64) {

	this.DataTotalCount = count
}
