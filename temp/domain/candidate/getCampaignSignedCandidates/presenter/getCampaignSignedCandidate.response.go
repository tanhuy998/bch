package presenter

import "app/src/model"

type GetCampaignSignedCandidatesResponse struct {
	Message              string               `json:"message"`
	Data                 []*model.Candidate   `json:"data"`
	DataTotalCount       int64                `json:"dataTotalCount"`
	CandidateSignedCount int64                `json:"candidateSignedCount"`
	Navigation           PaginationNavigation `json:"navigation"`
}

func (this *GetCampaignSignedCandidatesResponse) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}

func (this *GetCampaignSignedCandidatesResponse) SetTotalCount(count int64) {

	this.DataTotalCount = count
}
