package responsePresenter

import (
	"app/src/model"
)

type (
	GetCampaignListResponse struct {
		Message        string               `json:"message"`
		Data           []*model.Campaign    `json:"data"`
		DataTotalCount int64                `json:"dataTotalCount"`
		Navigation     PaginationNavigation `json:"navigation"`
	}
)

func (this *GetCampaignListResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}

func (this *GetCampaignListResponse) SetTotalCount(count int64) {

	this.DataTotalCount = count
}
