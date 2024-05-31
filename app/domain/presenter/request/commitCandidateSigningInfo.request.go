package requestPresenter

import "app/domain/model"

type CommitCandidateSigningInfoRequest struct {
	CandidateUUID string                      `param:"candidateUUID" validate:"required"`
	CampaignUUID  string                      `param:"campaignUUID" validate:"required"`
	Data          *model.CandidateSigningInfo `json:"data" validate:"required"`
}
