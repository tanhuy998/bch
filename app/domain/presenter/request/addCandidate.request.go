package requestPresenter

import "app/domain/model"

type AddCandidateRequest struct {
	CampaignUUID   string           `param:"campaignUUID" validate:"required"`
	InputCandidate *model.Candidate `json:"data" validate:"required"`
}
