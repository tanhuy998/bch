package adminService

import (
	"app/domain/model"
	"app/repository"
)

type (
	IGetCampaignCandidateList interface {
		Serve(campaignUUID string, candiatePivot_id string, limit int, isPrevDir bool) (*repository.PaginationPack[model.Candidate], error)
	}

	AdminGetCampaignCandidateListService struct {
		CampaignRepo  repository.ICampaignRepository
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *AdminGetCampaignCandidateListService) Serve(
	campaignUUID string, candiatePivot_id string, limit int, isPrevDir bool,
) (*repository.PaginationPack[model.Candidate], error) {

}
