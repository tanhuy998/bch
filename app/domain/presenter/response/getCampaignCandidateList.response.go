package responsePresenter

import "app/domain/model"

type GetCampaignCandidateListResponse struct {
	Message string `json:"message"`
	Header  struct {
		Campaign *model.Campaign `json:"campaign,omitempty"`
	} `json:"header,omitempty"`
	Data       []*model.Candidate   `json:"data"`
	Navigation PaginationNavigation `json:"navigation"`
}

func (this *GetCampaignCandidateListResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}
