package requestPresenter

import "app/domain/model"

type AddCandidateRequest struct {
	CampaignUUID    string           `param:"", validate:"required,uuid_rfc4122"`
	CandidateDetail *model.Candidate `json:"candidate" validate:"required"`
}
