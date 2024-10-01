package responsePresenter

import "app/src/model"

type GetCampaignCandidateListResponse struct {
	Message string `json:"message"`
	Header  struct {
		Campaign *model.Campaign `json:"campaign,omitempty"`
	} `json:"header,omitempty"`
	Data           []*model.Candidate   `json:"data"`
	DataTotalCount int64                `json:"dataTotalCount"`
	Navigation     PaginationNavigation `json:"navigation"`
}

func (this *GetCampaignCandidateListResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}

func (this *GetCampaignCandidateListResponse) SetTotalCount(count int64) {

	this.DataTotalCount = count
}
