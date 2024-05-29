package responsePresenter

import "app/domain/model"

type GetCampaignCandidateListResponse struct {
	Message    string               `json:"message"`
	Data       []*model.Candidate   `json:"data"`
	Navigation PaginationNavigation `json:"navigation"`
}

func (this *GetCampaignCandidateListResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}
