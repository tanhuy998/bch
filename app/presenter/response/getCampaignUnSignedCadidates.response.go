package responsePresenter

import "app/model"

type (
	GetCampaignUnSignedCandidates struct {
		Message              string               `json:"message"`
		Data                 []*model.Candidate   `json:"data"`
		DataTotalCount       int64                `json:"dataTotalCount"`
		CandidateSignedCount int64                `json:"candidateSignedCount"`
		Navigation           PaginationNavigation `json:"navigation"`
	}
)

func (this *GetCampaignUnSignedCandidates) GetNavigation() *PaginationNavigation {

	return &this.Navigation
}

func (this *GetCampaignUnSignedCandidates) SetTotalCount(count int64) {

	this.DataTotalCount = count
}
